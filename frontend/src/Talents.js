import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Fetch, pct, Filter } from './common';
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
		if (!this.props.build) {
			if (search === '') {
				search = '?';
			} else {
				search += '&';
			}
			search += 'build=' + this.props.Builds[0].ID;
		}
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
		if (!this.state.winrates) {
			return 'loading...';
		}
		const heroes = this.props.Heroes.map(h => (
			<option key={h.Name}>{h.Name}</option>
		));
		const heroSearch = this.props.build
			? '?build=' + encodeURIComponent(this.props.build)
			: '';
		return (
			<div>
				<h2>{this.props.match.params.hero}</h2>
				<p>
					<Link
						to={
							'/heroes/' + encodeURI(this.props.match.params.hero) + heroSearch
						}
					>
						[relative winrates]
					</Link>&nbsp;
					<a href="#popular">[popular builds]</a>&nbsp;
					<a href="#winning">[winning builds]</a>
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
				<Builds winrates={this.state.winrates} />
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

const TalentImg = props => {
	const words = props.name.match(/[A-Z][a-z]+/g);
	words.shift();
	const name = words.join(' ');
	return [
		<img
			key={props.name}
			src={'/img/talent/' + words.join('')}
			alt={name}
			title={name}
			style={{ verticalAlign: 'middle', paddingRight: '2px' }}
		/>,
		props.text ? name : null,
	];
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
			<div key={tier}>
				Tier {tierNames[tier]}
				<SortedTable
					sort="talent"
					headers={[
						{
							name: 'talent',
							cell: v => <TalentImg name={v} text={true} />,
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
				/>
			</div>
		);
	}
	function buildList(builds) {
		if (!builds) {
			return;
		}
		const list = builds.map((b, i) => (
			<tr key={i}>
				<td>{b.Build.map(v => <TalentImg key={v} name={v} />)}</td>
				<td>{b.Total}</td>
				<td>{pct(b.Winrate * 100)}</td>
			</tr>
		));
		return (
			<table>
				<thead>
					<tr>
						<th>build</th>
						<th>games</th>
						<th>winrate</th>
					</tr>
				</thead>
				<tbody>{list}</tbody>
			</table>
		);
	}
	const popular = buildList(props.winrates.PopularBuilds);
	const winning = buildList(props.winrates.WinningBuilds);
	return [
		<div key="builds">{builds}</div>,
		<div key="popular">
			<div className="anchor" id="popular" />
			<h3>Popular Builds</h3>
			{popular}
		</div>,
		<div key="winning">
			<div className="anchor" id="winning" />
			<h3>Winning Builds</h3>
			{winning}
		</div>,
	];
};

export default TalentWinrates;