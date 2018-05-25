// @flow

import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import {
	Fetch,
	pct,
	toLength,
	toDate,
	TalentImg,
	BuildsOpts,
	createCookie,
	readCookie,
	regionCookie,
	skillBands,
	skillPercentiles,
	regionNames,
	HeroImg,
} from './common';
import { Helmet } from 'react-helmet';
import SortedTable from './SortedTable';

import '../node_modules/react-vis/dist/style.css';
import {
	XYPlot,
	XAxis,
	YAxis,
	LineMarkSeries,
	DiscreteColorLegend,
} from 'react-vis';

class MMRChart extends Component<
	{
		xitems: any,
		series: any,
	},
	{}
> {
	textStyle = { text: { fill: '#fff' } };
	colors = {
		'Quick Match': '#4285f4',
		'Unranked Draft': '#1b9e77',
		'Hero League': '#db4437',
		'Team League': '#f4b400',
	};
	render() {
		return (
			<span style={{ display: 'inline-flex' }}>
				<XYPlot width={500} height={250}>
					<XAxis
						hideLine
						tickFormat={v => this.props.xitems[v]}
						tickLabelAngle={-90}
						tickSize={0}
						style={this.textStyle}
					/>
					<YAxis
						hideLine
						tickFormat={v => v + '%'}
						tickSize={2}
						style={this.textStyle}
						title="quantile"
					/>
					{this.props.series.map((v, i) => (
						<LineMarkSeries
							key={v.name}
							data={v.data}
							lineStyle={{ stroke: this.colors[v.name] }}
							markStyle={{
								stroke: this.colors[v.name],
								fill: this.colors[v.name],
							}}
							size={3}
						/>
					))}
				</XYPlot>
				<DiscreteColorLegend
					items={this.props.series.map(v => {
						return { title: v.name, color: this.colors[v.name] };
					})}
					style={this.textStyle.text}
				/>
			</span>
		);
	}
}

type Props = {
	match: any,
};

class Players extends Component<
	Props,
	{
		name: string,
		ids: ?(any[]),
		loading: boolean,
		region: string,
	}
> {
	constructor(props: Props) {
		super(props);
		this.state = {
			name: '',
			region: readCookie(regionCookie) || '1',
			ids: null,
			loading: false,
		};
	}
	update = (ev: SyntheticInputEvent<HTMLInputElement>) => {
		const st = {};
		st[ev.target.name] = ev.target.value;
		if (ev.target.name === 'region') {
			createCookie(regionCookie, ev.target.value);
		}
		this.setState(st);
	};
	onSubmit = (ev: SyntheticInputEvent<HTMLFormElement>) => {
		ev.preventDefault();
		if (!this.state.name) {
			return;
		}
		const name = this.state.name;
		this.setState({ loading: true });
		Fetch(
			'/api/get-player-by-name?name=' +
				encodeURIComponent(this.state.name) +
				'&region=' +
				this.state.region,
			data => {
				if (this.state.name === name) {
					this.setState({ ids: data, loading: false });
				}
			}
		);
	};
	render() {
		let names;
		if (this.state.loading) {
			names = 'loading...';
		} else if (!this.state.ids) {
		} else if (this.state.ids.length === 0) {
			names = 'No matches.';
		} else {
			names = (
				<SortedTable
					name="players"
					sort="Name"
					headers={[
						{
							name: 'Name',
							header: 'battletag',
							cell: (v, row) => (
								<Link to={'/players/' + row.Region + '/' + row.ID}>{v}</Link>
							),
						},
						{
							name: 'Games',
							header: 'games',
						},
					]}
					data={this.state.ids}
				/>
			);
		}
		return (
			<div>
				<Helmet>
					<title>Search for player by name</title>
				</Helmet>
				<form onSubmit={this.onSubmit}>
					<div className="row">
						<div className="column">
							<label>Search for player by battletag</label>
							<input
								type="text"
								name="name"
								value={this.state.name}
								onChange={this.update}
							/>
						</div>
						<div className="column">
							<label>Region</label>
							<select
								name="region"
								value={this.state.region}
								onChange={this.update}
							>
								{Object.keys(regionNames)
									.sort()
									.map(k => (
										<option key={k} value={k}>
											{regionNames[k]}
										</option>
									))}
							</select>
						</div>
					</div>
					<input className="button-primary" type="submit" value="search" />
				</form>
				{names}
			</div>
		);
	}
}

class Player extends Component<
	{
		match: any,
		Modes: any,
		Builds: any,
		build?: string,
		handleChange: any,
		history: any,
		HeroMap: any,
	},
	{
		Profile: any,
		Battletag: string,
		search: string,
		Skills: any,
		AllSkills: any,
		BuildStats?: any,
	}
> {
	state = {
		Profile: null,
		Battletag: '',
		search: '',
		Skills: [],
		AllSkills: [],
		BuildStats: null,
	};
	componentDidMount() {
		this.update();
	}
	componentDidUpdate() {
		this.update();
	}
	makeSearch = () => {
		const build = this.props.build || '90';
		return (
			'/api/get-player-profile?blizzid=' +
			encodeURIComponent(this.props.match.params.id) +
			'&region=' +
			this.props.match.params.region +
			'&build=' +
			encodeURIComponent(build)
		);
	};
	update = () => {
		const search = this.makeSearch();
		if (this.state.search === search) {
			return;
		}
		this.setState({ Profile: null, search: search });
		Fetch(search, data => {
			if (search === this.makeSearch()) {
				this.setState(data);
			}
		});
	};
	makeTable(name: string, prop: string, displayFn?: any) {
		const obj = this.state.Profile ? this.state.Profile[prop] : {};
		const elems = Object.keys(obj).map(k => {
			const v = obj[k];
			const total = v.Wins + v.Losses;
			const wr = v.Wins / total * 100;
			return {
				header: k,
				games: total,
				winrate: wr,
			};
		});
		return (
			<SortedTable
				name="profile"
				sort="header"
				headers={[
					{
						name: 'header',
						cell: displayFn,
						header: [
							<div key="anchor" className="anchor" id={prop.toLowerCase()} />,
							<span key="name">{name.toLowerCase()}</span>,
						],
					},
					{
						name: 'games',
						desc: true,
					},
					{
						name: 'winrate',
						cell: pct,
						desc: true,
					},
				]}
				data={elems}
				notable={true}
			/>
		);
	}
	percentile(
		skill: number,
		mode: number,
		stats: any
	): { rank: string, quantile: number } {
		if (!stats) {
			stats = this.state.BuildStats && this.state.BuildStats[mode];
		}
		if (!stats) {
			return { rank: 'unknown', quantile: 0 };
		}
		const qs = Object.keys(stats.Quantile);
		qs.sort((a, b) => a - b);
		if (skill < stats.Quantile[qs[0]]) {
			return { rank: skillBands[0], quantile: 0 };
		}
		for (let i = 1; i < qs.length; i++) {
			const q1 = qs[i];
			const p1 = stats.Quantile[q1];
			if (p1 >= skill) {
				const q0 = qs[i - 1];
				const p0 = stats.Quantile[q0];
				// If p lies a fraction f of the way from p{i} to p{i+1}, define the pth
				// quantile to be:
				// Q(p) = (1-f)Q(p{i}) + fQ(p{i+1})
				const f = (skill - p0) / (p1 - p0);
				const qp = (1 - f) * q0 + f * q1;
				return { rank: skillBands[i - 1], quantile: qp };
			}
		}
		return { rank: skillBands[skillBands.length - 1], quantile: 100 };
	}
	render() {
		let content, skillChart;
		if (!this.state.Profile) {
			content = 'loading...';
		} else {
			let skills;
			if (this.state.AllSkills) {
				const buildSet = {};
				this.state.AllSkills.forEach(v => {
					buildSet[v.Build] = true;
				});
				const builds = Object.keys(buildSet).sort();
				const series = [];
				Object.keys(this.props.Modes).forEach(modeID => {
					modeID = +modeID;
					const data = [];
					this.state.AllSkills.forEach(v => {
						if (v.Mode === modeID && v.Stats.Quantile) {
							const percentile = this.percentile(v.Skill, 0, v.Stats);
							data.push({ x: builds.indexOf(v.Build), y: percentile.quantile });
						}
					});
					if (data.length) {
						series.push({ data: data, name: this.props.Modes[modeID] });
					}
				});
				skillChart = <MMRChart xitems={builds} series={series} />;
			}
			if (this.state.Skills && this.state.BuildStats) {
				skills = (
					<SortedTable
						name="rank"
						sort="Mode"
						headers={[
							{
								header: 'game mode',
								name: 'Mode',
								cell: v => this.props.Modes[v],
							},
							{
								name: 'rank',
								title: skillBands
									.map(
										(v, i) =>
											skillPercentiles[i] +
											'-' +
											skillPercentiles[i + 1] +
											': ' +
											v
									)
									.join(', '),
							},
							{
								name: 'quantile',
								cell: pct,
							},
							{
								header: 'skill',
								name: 'Skill',
								cell: v => v.toFixed(6),
								title: 'based on trueskill (initial mean: 25.0)',
							},
						]}
						data={this.state.Skills.map(v =>
							Object.assign({}, v, this.percentile(v.Skill, v.Mode))
						)}
					/>
				);
			}
			content = (
				<div>
					{skills}
					<table className="sorted">
						{this.makeTable('Game Mode', 'Modes', m => this.props.Modes[m])}
						{this.makeTable('Role', 'Roles')}
						{this.makeTable('Hero', 'Heroes', v => (
							<HeroImg name={v} slug={this.props.HeroMap[v].Slug} />
						))}
						{this.makeTable('Map', 'Maps')}
					</table>
				</div>
			);
		}
		return (
			<div>
				<PlayerHeader
					{...this.props.match.params}
					history={this.props.history}
					name={this.state.Battletag}
					build={this.props.build}
					prefix="profile"
				/>
				<p>
					<a href="#modes">[game modes]</a> <a href="#roles">[roles]</a>{' '}
					<a href="#heroes">[heroes]</a> <a href="#maps">[maps]</a>
				</p>
				{skillChart}
				<PlayerOpts {...this.props} />
				{content}
			</div>
		);
	}
}

const PlayerOpts = (props: { Builds: any, handleChange: any, build: any }) => (
	<div className="row">
		<div className="column">
			<label>Patch</label>
			<select name="build" value={props.build} onChange={props.handleChange}>
				<BuildsOpts builds={props.Builds} dates={[30, 90]} />
			</select>
		</div>
		<div className="column" />
	</div>
);

class PlayerGames extends Component<
	{
		match: any,
		Modes: any,
		history: any,
		build?: string,
		handleChange: any,
		Builds: any,
	},
	{ Games: any[] | null, Battletag: string, search: string }
> {
	state = {
		Games: null,
		Battletag: '',
		search: '',
	};
	componentDidMount() {
		this.update();
	}
	componentDidUpdate() {
		this.update();
	}
	makeSearch = () => {
		const build = this.props.build || '90';
		return (
			'/api/get-player-games?blizzid=' +
			encodeURIComponent(this.props.match.params.id) +
			'&region=' +
			this.props.match.params.region +
			'&build=' +
			encodeURIComponent(build)
		);
	};
	update = () => {
		const search = this.makeSearch();
		if (this.state.search === search) {
			return;
		}
		this.setState({ Games: null, search: search });
		Fetch(search, data => {
			if (search === this.makeSearch()) {
				data.Games = data.Games || [];
				this.setState(data);
			}
		});
	};
	render() {
		let content;
		if (this.state.Games === null) {
			content = 'loading...';
		} else {
			content = (
				<SortedTable
					name="player"
					sort="Date"
					headers={[
						{
							name: 'Build',
							header: 'Patch',
							desc: true,
						},
						{
							name: 'Date',
							cell: (v, row) => (
								<Link to={'/games/' + row.Game}>{toDate(v)}</Link>
							),
							desc: true,
						},
						{
							name: 'Hero',
							cell: v => <Link to={'/talents/' + encodeURI(v)}>{v}</Link>,
						},
						{
							name: 'HeroLevel',
							header: 'Level',
							title: 'hero level',
							desc: true,
						},
						{
							name: 'Winner',
							header: 'Won',
							cell: v => (v ? 'win' : 'loss'),
							desc: true,
						},
						{
							name: 'Length',
							title: 'game length',
							cell: toLength,
						},
						{
							name: 'Map',
						},
						{
							name: 'Mode',
							header: 'Game Mode',
							cell: v => this.props.Modes[v],
						},
					]}
					data={this.state.Games}
				/>
			);
		}
		return (
			<div>
				<PlayerHeader
					{...this.props.match.params}
					history={this.props.history}
					name={this.state.Battletag}
					build={this.props.build}
					prefix="games"
				/>
				<PlayerOpts {...this.props} />
				{content}
			</div>
		);
	}
}

class PlayerMatchups extends Component<
	{
		match: any,
		Modes: any,
		history: any,
		build?: string,
		handleChange: any,
		Builds: any,
		HeroMap: any,
	},
	{ Same: any, Opposing: any, Battletag: string, search: string }
> {
	state = {
		Same: null,
		Opposing: null,
		Battletag: '',
		search: '',
	};
	componentDidMount() {
		this.update();
	}
	componentDidUpdate() {
		this.update();
	}
	makeSearch = () => {
		const build = this.props.build || '90';
		return (
			'/api/get-player-matchups?blizzid=' +
			encodeURIComponent(this.props.match.params.id) +
			'&region=' +
			this.props.match.params.region +
			'&build=' +
			encodeURIComponent(build)
		);
	};
	update = () => {
		const search = this.makeSearch();
		if (this.state.search === search) {
			return;
		}
		this.setState({ Same: null, search: search });
		Fetch(search, data => {
			if (search === this.makeSearch()) {
				this.setState(data);
			}
		});
	};
	makeTable(name: string, prop: string) {
		const obj = this.state[prop] ? this.state[prop] : {};
		const elems = Object.keys(obj).map(k => {
			const v = obj[k];
			const total = v.Wins + v.Losses;
			const wr = v.Wins / total * 100;
			return {
				header: k,
				games: total,
				winrate: wr,
			};
		});
		return (
			<SortedTable
				name="profile"
				sort="header"
				headers={[
					{
						name: 'header',
						cell: v => <HeroImg name={v} slug={this.props.HeroMap[v].Slug} />,
						header: [
							<div key="anchor" className="anchor" id={prop.toLowerCase()} />,
							<span key="name">{name.toLowerCase()}</span>,
						],
					},
					{
						name: 'games',
						desc: true,
					},
					{
						name: 'winrate',
						cell: pct,
						desc: true,
						title: 'winrate of ' + this.state.Battletag,
					},
				]}
				data={elems}
				notable={true}
			/>
		);
	}
	render() {
		let content;
		if (!this.state.Same) {
			content = 'loading...';
		} else {
			content = (
				<div>
					<table className="sorted">
						{this.makeTable('Opposing Team', 'Opposing')}
						{this.makeTable('Same Team', 'Same')}
					</table>
				</div>
			);
		}
		return (
			<div>
				<PlayerHeader
					{...this.props.match.params}
					history={this.props.history}
					name={this.state.Battletag}
					build={this.props.build}
					prefix="matchups/duos"
				/>
				<p>
					<a href="#opposing">[opposing team]</a>{' '}
					<a href="#same">[same team]</a>
				</p>
				<PlayerOpts {...this.props} />
				{content}
			</div>
		);
	}
}

class PlayerFriends extends Component<
	{
		match: any,
		history: any,
		build?: string,
	},
	{ Battletag: string, Friends: any, Region: string }
> {
	state = {
		Battletag: '',
		Friends: null,
		Region: '',
	};
	componentDidMount() {
		const url =
			'/api/get-player-friends?blizzid=' +
			encodeURIComponent(this.props.match.params.id) +
			'&region=' +
			this.props.match.params.region;
		Fetch(url, data => {
			this.setState(data);
		});
	}
	render() {
		let content;
		if (!this.state.Friends) {
			content = 'loading...';
		} else {
			content = (
				<SortedTable
					name="friends"
					sort="Games"
					headers={[
						{
							name: 'Battletag',
							header: 'name',
							cell: (v, row) => (
								<Link to={'/players/' + this.state.Region + '/' + row.Blizzid}>
									{v}
								</Link>
							),
						},
						{
							name: 'Games',
							header: 'games',
							desc: true,
						},
						{
							name: 'Winrate',
							header: 'winrate',
							cell: pct,
						},
					]}
					data={this.state.Friends}
				/>
			);
		}
		return (
			<div>
				<PlayerHeader
					{...this.props.match.params}
					history={this.props.history}
					name={this.state.Battletag}
					build={this.props.build}
					prefix="friends"
				/>
				{content}
			</div>
		);
	}
}

class Game extends Component<
	{ match: any, Modes: any, HeroMap: any },
	{ Game: any, Talents: any, Players: any }
> {
	componentDidMount() {
		Fetch(
			'/api/get-game-data?id=' + encodeURIComponent(this.props.match.params.id),
			data => {
				this.setState(data);
			}
		);
	}
	render() {
		if (!this.state) {
			return 'loading...';
		}
		const baseSearch = '?build=' + this.state.Game.Build;
		const title = this.state.Game.Map + ' on ' + toDate(this.state.Game.Date);
		const allColumns = {};
		this.state.Players.map(v => v.Data).forEach(v =>
			Object.keys(v).forEach(k => (allColumns[k] = true))
		);
		const columns = Object.keys(allColumns).sort();
		const dataHeaders = columns.map(v => ({
			name: v,
			key: v,
			header: v.replace(/_/g, ' '),
			cell: d => (d && v.startsWith('time') ? toLength(d) : d),
		}));
		dataHeaders.unshift({
			name: 'Hero',
			cell: v => <Link to={'/talents/' + v + baseSearch}>{v}</Link>,
		});
		return (
			<div>
				<Helmet>
					<title>{title}</title>
				</Helmet>
				<h2>{title}</h2>
				<table>
					<tbody>
						<tr>
							<td>Game Mode</td>
							<td>
								<Link to={'/' + baseSearch + '&mode=' + this.state.Game.Mode}>
									{this.props.Modes[this.state.Game.Mode]}
								</Link>
							</td>
						</tr>
						<tr>
							<td>Patch</td>
							<td>
								<Link to={'/' + baseSearch}>{this.state.Game.Build}</Link>
							</td>
						</tr>
						<tr>
							<td>Length</td>
							<td>{toLength(this.state.Game.Length)}</td>
						</tr>
						<tr>
							<td>Map</td>
							<td>
								<Link
									to={
										'/' +
										baseSearch +
										'&map=' +
										encodeURIComponent(this.state.Game.Map)
									}
								>
									{this.state.Game.Map}
								</Link>
							</td>
						</tr>
						<tr>
							<td>Date</td>
							<td>{toDate(this.state.Game.Date)}</td>
						</tr>
						<tr>
							<td>Bans</td>
							<td>{this.state.Game.BanList.join(', ')}</td>
						</tr>
					</tbody>
				</table>
				<SortedTable
					name="game"
					sort="Hero"
					headers={[
						{
							name: 'Hero',
							cell: v => (
								<HeroImg
									name={v}
									slug={this.props.HeroMap[v].Slug}
									link={baseSearch}
								/>
							),
						},
						{
							name: 'HeroLevel',
							header: 'Level',
							title: 'hero level',
							desc: true,
						},
						{
							name: 'Winner',
							header: 'Won',
							cell: v => (v ? 'win' : 'loss'),
							desc: true,
						},
						{
							name: 'Battletag',
							cell: (v, row) => (
								<Link to={'/players/' + row.Region + '/' + row.Blizzid}>
									{v}
								</Link>
							),
						},
						{
							name: 'TalentList',
							header: 'Talents',
							cell: v => (
								<span style={{ whiteSpace: 'nowrap' }}>
									{v.map(v => (
										<TalentImg key={v} name={v} data={this.state.Talents[v]} />
									))}
								</span>
							),
						},
					]}
					data={this.state.Players}
				/>
				<div style={{ overflowX: 'scroll' }}>
					<SortedTable
						name="data"
						sort="Hero"
						headers={dataHeaders}
						data={this.state.Players.map(v =>
							Object.assign({ Hero: v.Hero }, v.Data)
						)}
					/>
				</div>
			</div>
		);
	}
}

const PlayerHeader = (props: {
	name: string,
	id: string,
	region: string,
	history: any,
	build?: string,
	prefix: string,
}) => {
	const getClass = s =>
		props.history.location.pathname.endsWith(s)
			? 'button button-outline'
			: 'button';
	const build = props.build ? '?build=' + encodeURIComponent(props.build) : '';
	const title = props.name + ' ' + props.prefix;
	const link = '/players/' + props.region + '/' + props.id;
	return (
		<div>
			<h2>{title}</h2>
			<Helmet>
				<title>{title}</title>
			</Helmet>
			<Link key="profile" className={getClass(props.id)} to={link + build}>
				profile
			</Link>{' '}
			<Link
				key="games"
				className={getClass('games')}
				to={link + '/games' + build}
			>
				games
			</Link>{' '}
			<Link
				key="matchups"
				className={getClass('matchups')}
				to={link + '/matchups' + build}
			>
				matchups/duos
			</Link>{' '}
			<Link className={getClass('friends')} to={link + '/friends'}>
				friends
			</Link>
		</div>
	);
};

export { Players, Player, Game, PlayerGames, PlayerMatchups, PlayerFriends };
