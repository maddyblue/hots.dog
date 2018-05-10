import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Fetch, pct, Filter, HeroHeader } from './common';
import SortedTable from './SortedTable';

class CompareHero extends Component {
	constructor(props) {
		super(props);
		this.state = {};
		this.changeHero = this.changeHero.bind(this);
		this.makeSearch = this.makeSearch.bind(this);
	}
	changeHero(ev) {
		this.props.history.push({
			pathname: '/compare/' + encodeURI(ev.target.value),
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
			compare: null,
			search: search,
		});
		Fetch('/api/get-compare-hero' + search, data => {
			if (search === this.makeSearch()) {
				if (!data.Previous) {
					data.Previous = {};
				}
				this.setState({
					compare: data,
				});
			}
		});
	}
	render() {
		let builds;
		if (!this.state.compare) {
			builds = 'loading...';
		} else {
			const basetotal =
				this.state.compare.Total.Wins + this.state.compare.Total.Losses;
			const basewr = this.state.compare.Total.Wins / basetotal * 100;
			const sameTeam = [];
			const otherTeam = [];
			this.props.Heroes.forEach(hero => {
				const st = this.state.compare.SameTeam[hero.Name];
				if (st) {
					const games = st.Wins + st.Losses;
					const winrate = st.Wins / games * 100;
					sameTeam.push({
						hero: hero,
						games: games,
						winrate: winrate,
						relative: winrate - basewr,
					});
				}
				const ot = this.state.compare.OtherTeam[hero.Name];
				if (ot) {
					const games = ot.Wins + ot.Losses;
					const winrate = ot.Wins / games * 100;
					otherTeam.push({
						hero: hero,
						games: games,
						winrate: winrate,
						relative: winrate - basewr,
					});
				}
			});
			builds = [
				<table key="base">
					<thead>
						<tr>
							<th>games</th>
							<th>base winrate</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>{basetotal}</td>
							<td>{pct(basewr)}</td>
						</tr>
					</tbody>
				</table>,
				<div key="other">
					<div className="anchor" id="other" />
					<h2>Opposing Team</h2>
					<Winrates
						winrates={otherTeam}
						name="against hero"
						hero={this.props.match.params.hero}
						basewr={basewr}
						search={this.props.history.location.search}
					/>
				</div>,
				<div key="same">
					<div className="anchor" id="same" />
					<h2>Same Team</h2>
					<Winrates
						winrates={sameTeam}
						name="with hero"
						hero={this.props.match.params.hero}
						basewr={basewr}
						search={this.props.history.location.search}
					/>
				</div>,
			];
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
					prefix="comparison"
					title={this.props}
				/>
				<p>
					<a href="#other">[opposing team]</a> <a href="#same">[same team]</a>
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

const Winrates = props => (
	<SortedTable
		name="compare"
		sort="winrate"
		headers={[
			{
				name: 'hero',
				header: props.name,
				cmp: (a, b) => a.Name.localeCompare(b.Name),
				cell: v => (
					<Link to={'/compare/' + encodeURI(v.Name) + props.search}>
						<img
							key="img"
							src={'/img/hero/' + v.Slug + '.png'}
							alt={v.Name}
							style={{
								width: '40px',
								height: '40px',
								verticalAlign: 'middle',
								marginRight: '1em',
							}}
						/>
						{v.Name}
					</Link>
				),
			},
			{
				name: 'games',
				desc: true,
			},
			{
				name: 'winrate',
				title: 'winrate of ' + props.hero,
				cell: pct,
				desc: true,
			},
			{
				name: 'relative',
				title: props.hero + ' winrate relative to base winrate',
				cell: pct,
				desc: true,
			},
		]}
		data={props.winrates}
	/>
);

export default CompareHero;
