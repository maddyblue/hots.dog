import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Fetch, pct, readCookie, Filter } from './common';

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
	constructor(props) {
		super(props);
		this.state = {
			sort: 'winrate',
			sortDir: true,
		};
		this.sort = this.sort.bind(this);
		this.sortClass = this.sortClass.bind(this);
	}
	sort(ev) {
		const sort = ev.target.innerText;
		let dir;
		if (this.state.sort === sort) {
			dir = !this.state.sortDir;
		} else {
			switch (sort) {
				case 'hero':
					dir = false;
					break;
				case 'games':
				case 'winrate':
				case 'change':
					dir = true;
					break;
				default:
					console.log('unknown sort target:', sort);
					return;
			}
		}
		this.setState({ sort: sort, sortDir: dir });
	}
	sortClass(col) {
		if (col !== this.state.sort) {
			return '';
		}
		const dir = this.state.sortDir ? 'asc' : 'desc';
		return 'sort-' + dir;
	}
	render() {
		const sortedWinrates = this.props.winrates.concat();
		sortedWinrates.sort((a, b) => {
			switch (this.state.sort) {
				case 'hero':
					return a.hero.Name.localeCompare(b.hero.Name);
				case 'games':
				case 'winrate':
				case 'change':
					const as = a[this.state.sort];
					const bs = b[this.state.sort];
					return as - bs;
				default:
					console.log('unknown sort:', this.state.sort);
					return 0;
			}
		});
		if (this.state.sortDir) {
			sortedWinrates.reverse();
		}
		const build = this.props.build
			? '?build=' + encodeURIComponent(this.props.build)
			: '';
		const winrates = sortedWinrates.map(wr => {
			return (
				<tr key={wr.hero.Name}>
					<td>
						<img
							src={wr.hero.Icon}
							alt={wr.hero.Name}
							style={{
								width: '40px',
								height: '40px',
								verticalAlign: 'middle',
								marginRight: '1em',
							}}
						/>
						<Link to={'/talents/' + encodeURI(wr.hero.Name) + build}>
							{wr.hero.Name}
						</Link>
					</td>
					<td>{wr.games || 0}</td>
					<td>{pct(wr.winrate)}</td>
					<td>{pct(wr.change)}</td>
				</tr>
			);
		});
		return (
			<table>
				<thead>
					<tr>
						<th onClick={this.sort} className={this.sortClass('hero')}>
							hero
						</th>
						<th onClick={this.sort} className={this.sortClass('games')}>
							games
						</th>
						<th onClick={this.sort} className={this.sortClass('winrate')}>
							winrate
						</th>
						<th
							onClick={this.sort}
							className={this.sortClass('change')}
							title="percent change from previous patch"
						>
							change
						</th>
					</tr>
				</thead>
				<tbody>{winrates}</tbody>
			</table>
		);
	}
}

export default HeroWinrates;
