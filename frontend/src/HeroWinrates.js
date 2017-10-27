import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Fetch, pct, readCookie, Filter } from './common';
import SortedTable from './SortedTable';

class HeroWinrates extends Component {
	constructor(props) {
		super(props);
		this.state = {};
		const params = readCookie('params');
		if (params && !this.props.history.location.search) {
			this.props.history.replace({ search: params });
		}
	}
	componentDidMount() {
		this.update();
	}
	componentDidUpdate(prevProps, prevState) {
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
			winrates = <div>loading...</div>;
		} else {
			const rates = [];
			this.props.Heroes.forEach(hero => {
				const wr = this.state.winrates.Current[hero.Name];
				if (!wr) {
					return;
				}
				let games = 0;
				let winrate = 0;
				let change = 0;
				if (wr) {
					games = wr.Wins + wr.Losses;
					winrate = wr.Wins / games * 100;
					const prev =
						this.state.winrates.Previous &&
						this.state.winrates.Previous[hero.Name];
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
			winrates = <Winrates winrates={rates} build={this.props.build} />;
		}
		return (
			<div>
				<Filter handleChange={this.props.handleChange} {...this.props} />
				{winrates}
			</div>
		);
	}
}

class Winrates extends Component {
	render() {
		const build = this.props.build
			? '?build=' + encodeURIComponent(this.props.build)
			: '';
		return (
			<SortedTable
				sort="winrate"
				headers={[
					{
						name: 'hero',
						cmp: (a, b) => a.Name.localeCompare(b.Name),
						cell: v => (
							<Link to={'/talents/' + encodeURI(v.Name) + build} key="link">
								<img
									key="img"
									src={v.Icon}
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
				data={this.props.winrates}
			/>
		);
	}
}

export default HeroWinrates;
