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
	skillBands,
} from './common';
import { Helmet } from 'react-helmet';
import SortedTable from './SortedTable';

type Props = {
	match: any,
};

const regionNames = {
	'1': 'Americas',
	'2': 'Europe',
	'3': 'Asia',
	'5': 'China',
};

const regionCookie = 'region';

class Players extends Component<
	Props,
	{
		name: string,
		ids: any[],
		loading: boolean,
		region: string,
	}
> {
	constructor(props: Props) {
		super(props);
		this.state = {
			name: '',
			region: readCookie(regionCookie) || '1',
			ids: [],
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
			names = 'No matches.';
		} else {
			names = this.state.ids.map(e => (
				<li key={e.Name}>
					<Link to={'/players/' + e.Region + '/' + e.ID}>{e.Name}</Link>
				</li>
			));
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
				<ul>{names}</ul>
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
	},
	{
		Profile: any,
		Battletag: string,
		search: string,
		Skills?: any,
		SkillMult: number,
		BuildStats?: any,
	}
> {
	state = {
		Profile: null,
		Battletag: '',
		search: '',
		Skills: null,
		SkillMult: 0,
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
	percentile(skill: number, mode: number) {
		const stats = this.state.BuildStats && this.state.BuildStats[mode];
		if (!stats) {
			return 'unknown';
		}
		const qs = Object.keys(stats.Quantile);
		qs.sort((a, b) => a - b);
		if (skill < stats.Quantile[qs[0]]) {
			return skillBands[0];
			//return pct(0, 0);
		}
		for (let i = 1; i < qs.length; i++) {
			const q1 = qs[i];
			const p1 = stats.Quantile[q1];
			if (p1 >= skill) {
				return skillBands[i - 1];
				/*
				const q0 = qs[i - 1];
				const p0 = stats.Quantile[q0];
				// If p lies a fraction f of the way from p{i} to p{i+1}, define the pth
				// quantile to be:
				// Q(p) = (1-f)Q(p{i}) + fQ(p{i+1})
				const f = (skill - p0) / (p1 - p0);
				const qp = (1 - f) * q0 + f * q1;
				return pct(qp, 0);
				*/
			}
		}
		return skillBands[skillBands.length - 1];
		//return pct(100, 0);
	}
	render() {
		let content;
		if (!this.state.Profile) {
			content = 'loading...';
		} else {
			let skills;
			if (this.state.Skills && this.state.BuildStats) {
				skills = (
					<table>
						<tbody>
							{this.state.Skills.map(v => (
								<tr key={v.Mode}>
									<td>{this.props.Modes[v.Mode]}</td>
									<td>
										{this.percentile(v.Skill / this.state.SkillMult, v.Mode)}
									</td>
								</tr>
							))}
						</tbody>
					</table>
				);
			}
			content = (
				<div>
					{skills}
					<table className="sorted">
						{this.makeTable('Game Mode', 'Modes', m => this.props.Modes[m])}
						{this.makeTable('Role', 'Roles')}
						{this.makeTable('Hero', 'Heroes', v => (
							<Link to={'/talents/' + encodeURI(v)}>{v}</Link>
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
						cell: v => <Link to={'/talents/' + encodeURI(v)}>{v}</Link>,
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
					prefix="matchups"
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

class Game extends Component<
	{ match: any, Modes: any },
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
							cell: v => <Link to={'/talents/' + v + baseSearch}>{v}</Link>,
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
				matchups
			</Link>
		</div>
	);
};

export { Players, Player, Game, PlayerGames, PlayerMatchups };
