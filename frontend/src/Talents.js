import React, { Component } from 'react';
import { Fetch, pct, Filter, HeroHeader, TalentImg } from './common';
import SortedTable from './SortedTable';

class TalentWinrates extends Component {
	constructor(props) {
		super(props);
		this.state = {};
		this.changeHero = this.changeHero.bind(this);
		this.makeSearch = this.makeSearch.bind(this);
	}
	changeHero(ev) {
		this.props.history.push({
			pathname: '/talents/' + encodeURI(ev.target.value),
			search: this.props.history.location.search,
		});
	}
	componentDidMount() {
		this.update();
	}
	componentDidUpdate(prevProps, prevState) {
		this.update();
	}
	makeSearch() {
		let search = window.location.search;
		if (search) {
			search += '&';
		} else {
			search = '?';
		}
		search += 'hero=' + encodeURIComponent(this.props.match.params.hero);
		if (search.indexOf('build') !== -1) {
			return search;
		}
		if (search === '') {
			search = '?';
		} else {
			search += '&';
		}
		search += 'build=';
		search += this.props.build || this.props.Builds[0].ID;
		return search;
	}
	update() {
		const search = this.makeSearch();
		if (!search || search === this.state.search) {
			return;
		}
		this.setState({
			winrates: null,
			search: search,
		});
		Fetch('/api/get-build-winrates' + search, data => {
			if (search === this.makeSearch()) {
				if (!data.Previous) {
					data.Previous = {};
				}
				this.setState({
					winrates: data,
				});
			}
		});
	}
	render() {
		let builds;
		if (!this.state.winrates) {
			builds = 'loading...';
		} else {
			builds = <Builds winrates={this.state.winrates} />;
		}
		const heroes = this.props.Heroes.map(h => (
			<option key={h.Name}>{h.Name}</option>
		));
		return (
			<div>
				<HeroHeader
					name={this.props.match.params.hero}
					heroes={this.props.Heroes}
					build={this.props.build}
					history={this.props.history}
					prefix="talents"
					title={this.props}
				/>
				<p>
					<a href="#winning">[winning builds]</a>{' '}
					<a href="#popular">[popular builds]</a>
				</p>
				<div className="row">
					<div className="column">
						<label>Hero</label>
						<select
							name="hero"
							value={this.props.match.params.hero}
							onChange={this.changeHero}
						>
							{heroes}
						</select>
					</div>
					<div className="column" />
				</div>
				<Filter handleChange={this.props.handleChange} {...this.props} />
				{builds}
			</div>
		);
	}
}

const tierNames = {
	1: 1,
	2: 4,
	3: 7,
	4: 10,
	5: 13,
	6: 16,
	7: 20,
};

const Builds = props => {
	let builds = [];
	for (let tier = 1; tier <= 7; tier++) {
		const curTier = props.winrates.Current[tier];
		const prevTier = props.winrates.Previous[tier] || {};
		const talents = Object.keys(curTier).map(talent => {
			const cur = curTier[talent];
			const prev = prevTier[talent];
			const games = cur.Wins + cur.Losses;
			const winRate = cur.Wins / games * 100;
			let change = 0;
			if (prev) {
				const prevGames = prev.Wins + prev.Losses;
				const prevWinRate = prev.Wins / prevGames * 100;
				change = winRate - prevWinRate;
			}
			return {
				talent: talent,
				games: games,
				winrate: winRate,
				change: change,
			};
		});
		builds.push(
			<SortedTable
				key={tier}
				name="talent"
				sort="winrate"
				headers={[
					{
						name: 'talent',
						header: 'tier ' + tierNames[tier],
						cell: v => (
							<TalentImg
								name={v}
								text={true}
								data={props.winrates.Talents[v]}
							/>
						),
						cmp: (a, b) => {
							const x = props.winrates.Talents[a].Name || a;
							const y = props.winrates.Talents[b].Name || b;
							return x.localeCompare(y);
						},
					},
					{
						name: 'games',
						desc: true,
					},
					{
						name: 'winrate',
						desc: true,
						cell: pct,
					},
					{
						name: 'change',
						title: 'change since previous patch',
						cell: pct,
						desc: true,
					},
				]}
				data={talents}
				notable={true}
			/>
		);
	}
	function buildList(builds, sort, name) {
		if (!builds) {
			return;
		}
		return (
			<SortedTable
				name="popular"
				sort={sort}
				headers={[
					{
						name: 'Build',
						header: [
							<div key="anchor" className="anchor" id={name.toLowerCase()} />,
							<span key="name">{name} builds</span>,
						],
						cell: v =>
							v.map(v => (
								<TalentImg key={v} name={v} data={props.winrates.Talents[v]} />
							)),
						cmp: null,
					},
					{
						name: 'Total',
						header: 'games',
						desc: true,
					},
					{
						name: 'Winrate',
						header: 'winrate',
						cell: v => pct(v * 100),
						desc: true,
					},
				]}
				data={builds}
				notable={true}
			/>
		);
	}
	const winning = buildList(props.winrates.WinningBuilds, 'Winrate', 'winning');
	const popular = buildList(props.winrates.PopularBuilds, 'Total', 'popular');
	return [
		<div key="builds">
			<table className="sorted">{builds}</table>
		</div>,
		<table key="table" className="sorted">
			{winning}
			{popular}
		</table>,
	];
};

export default TalentWinrates;
