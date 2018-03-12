// @flow

import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Fetch, pct, Filter, Title } from './common';
import SortedTable from './SortedTable';

type Props = {
	handleChange: Event => void,
	map: string,
	build: string,
	herolevel: string,
	history: any,
	mode: string,
	skill_low: string,
	skill_high: string,
	Builds: any,
	Heroes: any[],
	Maps: string[],
	Modes: {},
	BuildStats: any,
};

class HeroWinrates extends Component<
	Props,
	{ winrates?: any, search?: string }
> {
	constructor(props: Props) {
		super(props);
		this.state = {};
	}
	componentDidMount() {
		this.update();
	}
	componentDidUpdate() {
		this.update();
	}
	makeSearch() {
		let search = window.location.search;
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
		Fetch('/api/get-winrates' + search, data => {
			if (search === this.makeSearch()) {
				this.setState({
					winrates: data,
					search: search,
				});
			}
		});
	}
	render() {
		let winrates;
		if (!this.state.winrates) {
			winrates = 'loading...';
		} else {
			const rates = [];
			const wrs = this.state.winrates || {};
			this.props.Heroes.forEach(hero => {
				const wr = wrs.Current[hero.Name];
				if (!wr) {
					return;
				}
				let games = 0;
				let winrate = 0;
				let change = 0;
				if (wr) {
					games = wr.Wins + wr.Losses;
					winrate = wr.Wins / games * 100;
					const prev = wrs.Previous && wrs.Previous[hero.Name];
					if (prev) {
						const prevGames = prev.Wins + prev.Losses;
						const prevWinRate = prev.Wins / prevGames * 100;
						change = winrate - prevWinRate;
					}
				}
				rates.push({
					hero: hero,
					games: games,
					winrate: winrate,
					change: change,
				});
			});
			winrates = rates.length ? (
				<Winrates winrates={rates} build={this.props.build} />
			) : (
				'No results. This could be because a map was out of HL/TL rotation, or there were no results with your filters.'
			);
		}
		return (
			<div>
				<Title {...this.props} prefix="Hero winrates" />
				<Filter handleChange={this.props.handleChange} {...this.props} />
				{winrates}
			</div>
		);
	}
}

const Winrates = props => {
	const build = props.build ? '?build=' + encodeURIComponent(props.build) : '';
	return (
		<SortedTable
			name="hero"
			sort="winrate"
			headers={[
				{
					name: 'hero',
					cmp: (a, b) => a.Name.localeCompare(b.Name),
					cell: v => (
						<Link to={'/talents/' + encodeURI(v.Name) + build} key="link">
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
					cell: v => v || 0,
					desc: true,
				},
				{
					name: 'winrate',
					cell: pct,
					desc: true,
				},
				{
					name: 'change',
					title: 'percent change from previous patch',
					cell: pct,
					desc: true,
				},
			]}
			data={props.winrates}
		/>
	);
};

export default HeroWinrates;
