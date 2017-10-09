import React, { Component } from 'react';
import { Link, Route, withRouter } from 'react-router-dom';
import './App.css';
import './normalize.css';
import './milligram.css';

const skillPercentiles = [0, 30, 50, 70, 90, 95, 100];

function Fetch(url, success, error) {
	if (!error) {
		error = alert;
	}
	fetch(url)
		.then(resp => {
			if (resp.status !== 200) {
				resp.text().then(error, error);
				return;
			}
			resp.json().then(success, error);
		})
		.catch(error);
}

class HotsApp extends Component {
	constructor(props) {
		super(props);
		this.params = [
			'build',
			'map',
			'mode',
			'herolevel',
			'skill_low',
			'skill_high',
		];
		this.defaultHeroLevel = '5';
		this.defaultSkillLow = '0';
		this.defaultSkillHigh = '100';
		this.state = this.searchState();
		this.handleChange = this.handleChange.bind(this);
	}
	componentDidMount() {
		Fetch('/api/init', data => {
			data.init = true;
			this.setState(data);
		});
	}
	componentDidUpdate(prevProps, prevState) {
		// This function attempts to sink the state with the URL when the URL changes,
		// for use with back/forward navigation actions.
		var locationChanged = this.props.location !== prevProps.location;
		if (!locationChanged) {
			return;
		}
		var st = this.searchState();
		var updates = {};
		this.params.forEach(key => {
			if (st[key] !== this.state[key]) {
				updates[key] = st[key] || '';
			}
		});
		if (Object.keys(updates).length === 0) {
			return;
		}
		this.setState(updates);
	}
	searchState() {
		const st = {};
		const search = new URLSearchParams(window.location.search);
		this.params.forEach(key => {
			if (!search.has(key)) {
				return;
			}
			st[key] = search.get(key);
		});
		if (!st.herolevel) {
			st.herolevel = this.defaultHeroLevel;
		}
		if (!st.skill_low) {
			st.skill_low = this.defaultSkillLow;
		}
		if (!st.skill_high) {
			st.skill_high = this.defaultSkillHigh;
		}
		return st;
	}
	getSearch() {
		const params = [];
		this.params.forEach(key => {
			const value = this.state[key];
			if (!value) {
				return;
			}
			if (key === 'herolevel' && value === this.defaultHeroLevel) {
				return;
			}
			if (
				key === 'skill_low' &&
				(value === this.defaultSkillLow || !this.state.mode)
			) {
				return;
			}
			if (
				key === 'skill_high' &&
				(value === this.defaultSkillHigh || !this.state.mode)
			) {
				return;
			}
			params.push(key + '=' + encodeURIComponent(value));
		});
		return '?' + params.join('&');
	}
	handleChange(event) {
		const st = {};
		st[event.target.name] = event.target.value;
		const ev = +event.target.value;
		switch (event.target.name) {
			case 'skill_low':
				if (ev >= +this.state.skill_high) {
					st['skill_high'] = skillPercentiles[skillPercentiles.indexOf(ev) + 1];
				}
				break;
			case 'skill_high':
				if (ev <= +this.state.skill_low) {
					st['skill_low'] = skillPercentiles[skillPercentiles.indexOf(ev) - 1];
				}
				break;
			default:
				break;
		}
		this.setState(st, () => {
			const params = this.getSearch();
			createCookie('params', params, 50);
			this.props.history.push({ search: params });
		});
	}
	render() {
		if (!this.state.init) {
			return 'loading...';
		}
		return (
			<main className="wrapper">
				<nav className="navigation">
					<section className="container">
						<Link to="/" className="navigation-title">
							home
						</Link>
						<ul className="navigation-list float-right">
							<li className="navigation-item">
								<Link className="navigation-title" to="/players">
									players
								</Link>
							</li>
							<li className="navigation-item">
								<Link className="navigation-link" to="/about">
									about
								</Link>
							</li>
						</ul>
					</section>
				</nav>
				<section className="container">
					<Route
						exact
						path="/"
						render={props => (
							<HeroWinrates
								params={this.params}
								handleChange={this.handleChange}
								{...this.state}
								{...props}
							/>
						)}
					/>
					<Route exact path="/about" component={About} />
					<Route exact path="/players" component={Players} />
					<Route
						exact
						path="/players/:id"
						render={props => <Player {...props} Modes={this.state.Modes} />}
					/>
				</section>
			</main>
		);
	}
}

const Filter = props => {
	let maps = props.Maps.map(m => <option key={m}>{m}</option>);
	maps.unshift(
		<option key="" value="">
			All Maps
		</option>
	);
	let builds = props.Builds.map(b => (
		<option key={b.ID} value={b.ID}>
			{b.ID} ({new Date(b.Start).toLocaleDateString()} -{' '}
			{new Date(b.Finish).toLocaleDateString()})
		</option>
	));
	builds.unshift(
		<option key="latest" value="">
			latest ({props.Builds[0].ID})
		</option>
	);
	let modeKeys = Object.keys(props.Modes);
	modeKeys.sort().reverse();
	let modes = modeKeys.map(k => (
		<option key={k} value={k}>
			{props.Modes[k]}
		</option>
	));
	modes.unshift(
		<option key="" value="">
			All Game Modes
		</option>
	);
	let heroLevels = [1, 5, 10, 20].map(v => <option key={v}>{v}</option>);
	const skills = skillPercentiles.map(v => (
		<option key={v} value={v}>
			{v + 'th'}
		</option>
	));
	const build = props.build === 'latest' ? props.Builds[0].ID : props.build;
	const buildStats = props.BuildStats[build];
	const modeStats = buildStats && buildStats[props.mode];
	const disableSkill = !modeStats || !props.mode;
	let skillTitle = 'Enabled when a game mode is selected.';
	if (!buildStats || (!modeStats && props.mode)) {
		skillTitle = 'Skill ratings not yet calculated.';
	}
	return (
		<div>
			<div className="row">
				<div className="column">
					<label>Map</label>
					<select name="map" value={props.map} onChange={props.handleChange}>
						{maps}
					</select>
				</div>
				<div className="column">
					<label>Patch</label>
					<select
						name="build"
						value={props.build}
						onChange={props.handleChange}
					>
						{builds}
					</select>
				</div>
			</div>
			<div className="row">
				<div className="column">
					<label>Game Mode</label>
					<select name="mode" value={props.mode} onChange={props.handleChange}>
						{modes}
					</select>
				</div>
				<div className="column">
					<label>Minimum Hero Level</label>
					<select
						name="herolevel"
						value={props.herolevel}
						onChange={props.handleChange}
					>
						{heroLevels}
					</select>
				</div>
			</div>
			<div className="row">
				<div className="column column-25">
					<label title={skillTitle}>Skill Percentile from</label>
					<select
						name="skill_low"
						value={props.skill_low}
						onChange={props.handleChange}
						disabled={disableSkill}
					>
						{skills.slice(0, -1)}
					</select>
				</div>
				<div className="column column-25">
					<label>to</label>
					<select
						name="skill_high"
						value={props.skill_high}
						onChange={props.handleChange}
						disabled={disableSkill}
					>
						{skills.slice(1)}
					</select>
				</div>
			</div>
		</div>
	);
};

const About = props => {
	return (
		<div>
			<h2>about</h2>
			<p>
				This site is a{' '}
				<a href="http://us.battle.net/heroes/en/">Heroes of the Storm</a> game
				aggregator. It fetches data from{' '}
				<a href="http://hotsapi.net/">HotsApi</a> and displays hero winrates
				with various filter options.
			</p>
			<p>Our goals are:</p>
			<ul>
				<li>a fast, clean web experience</li>
				<li>
					URL reflects current view; this means you can bookmark a specific
					filter view
				</li>
				<li>open API for others to use</li>
			</ul>
			<p>Our future goals are:</p>
			<ul>
				<li>data on talent picks</li>
				<li>player MMR calculation</li>
			</ul>
			<p>
				The code is free on GitHub at{' '}
				<a href="https://github.com/mjibson/hots-cockroach">
					github.com/mjibson/hots-cockroach
				</a>. Technical details:
			</p>
			<ul>
				<li>
					backend is written in <a href="https://golang.org/">Go</a>
				</li>
				<li>
					database is <a href="https://www.cockroachlabs.com/">CockroachDB</a>
				</li>
				<li>
					frontend is <a href="https://facebook.github.io/react/">React</a>
				</li>
				<li>
					site is deployed with <a href="https://kubernetes.io/">
						Kubernetes
					</a>{' '}
					on{' '}
					<a href="https://cloud.google.com/container-engine/">
						Google Container Engine
					</a>
				</li>
			</ul>
		</div>
	);
};

class Players extends Component {
	constructor(props) {
		super(props);
		this.state = {
			name: '',
			ids: [],
		};
		this.update = this.update.bind(this);
		this.onSubmit = this.onSubmit.bind(this);
	}
	update(ev) {
		const st = {};
		st[ev.target.name] = ev.target.value;
		this.setState(st);
	}
	onSubmit(ev) {
		ev.preventDefault();
		if (!this.state.name) {
			return;
		}
		Fetch(
			'/api/get-player-by-name?name=' + encodeURIComponent(this.state.name),
			data => this.setState({ ids: data })
		);
	}
	render() {
		const names = this.state.ids.map(e => (
			<li key={e.ID}>
				<Link to={'/players/' + e.ID}>
					{e.Name}: {e.ID}
				</Link>
			</li>
		));
		return (
			<div>
				<form onSubmit={this.onSubmit}>
					<label>Search for player by name</label>
					<input
						type="text"
						name="name"
						value={this.state.name}
						onChange={this.update}
					/>
					<input className="button-primary" type="submit" value="search" />
				</form>
				<ul>{names}</ul>
			</div>
		);
	}
}

class Player extends Component {
	componentDidMount() {
		Fetch(
			'/api/get-player-data?id=' +
				encodeURIComponent(this.props.match.params.id),
			data => {
				if (!data.Skills) {
					data.Skills = [];
				}
				this.setState(data);
			}
		);
	}
	render() {
		if (!this.state) {
			return 'loading...';
		}
		const skills = this.state.Skills.map((s, i) => (
			<tr key={i}>
				<td>{s.Build}</td>
				<td>{this.props.Modes[s.Mode]}</td>
				<td>{s.Skill}</td>
			</tr>
		));
		const games = this.state.Games.map((g, i) => (
			<tr key={i}>
				<td>{g.Build}</td>
				<td>{g.Hero}</td>
				<td>{g.HeroLevel}</td>
				<td>{g.Winner ? 'win' : 'loss'}</td>
				<td>{toLength(g.Length)}</td>
				<td>{g.Map}</td>
				<td>{this.props.Modes[g.Mode]}</td>
				<td>{g.Skill}</td>
			</tr>
		));
		return (
			<div>
				<p>Skill rating at the end of each patch with played games:</p>
				<table>
					<thead>
						<tr>
							<th>Patch</th>
							<th>Game Mode</th>
							<th>Skill Rating</th>
						</tr>
					</thead>
					<tbody>{skills}</tbody>
				</table>
				<table>
					<thead>
						<tr>
							<th>Patch</th>
							<th>Hero</th>
							<th title="Hero level">Level</th>
							<th>Won</th>
							<th title="Game length">Length</th>
							<th>Map</th>
							<th>Game Mode</th>
							<th>Skill Rating</th>
						</tr>
					</thead>
					<tbody>{games}</tbody>
				</table>
			</div>
		);
	}
}

class TalentWinrates extends Component {
	constructor(props) {
		super(props);
		this.state = {};
	}
	componentDidMount() {
		this.update();
	}
	componentDidUpdate(prevProps, prevState) {
		this.update();
	}
	update() {
		const search = window.location.search;
		if (search === this.state.search) {
			return;
		}
		Fetch(
			'/api/get-build-winrates' + this.props.match.params.hero + search,
			data => {
				if (search === window.location.search) {
					this.setState({
						winrates: data,
						search: search,
					});
				}
			}
		);
	}
	render() {
		if (!this.state.winrates) {
			return 'loading...';
		}
		return <Builds winrates={this.state.winrates} />;
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
	var builds = [];
	for (let tier = 1; tier <= 7; tier++) {
		const curTier = props.winrates.Current[tier];
		const prevTier = props.winrates.Previous[tier];
		const talents = Object.keys(curTier)
			.sort()
			.map(talent => {
				const cur = curTier[talent];
				const prev = prevTier[talent];
				const games = cur.Wins + cur.Losses;
				const winRate = cur.Wins / games * 100;
				const rate = winRate.toFixed(1) + '%';
				let change;
				if (prev) {
					const prevGames = prev.Wins + prev.Losses;
					const prevWinRate = prev.Wins / prevGames * 100;
					change = (winRate - prevWinRate).toFixed(1) + '%';
				}
				return (
					<tr key={talent}>
						<td>{talent}</td>
						<td>{games}</td>
						<td>{rate}</td>
						<td>{change}</td>
					</tr>
				);
			});
		builds.push(
			<div key={tier}>
				Tier {tierNames[tier]}
				<table>
					<thead>
						<tr>
							<th>hero</th>
							<th>games</th>
							<th>winrate</th>
							<th title="change since previous patch">change</th>
						</tr>
					</thead>
					<tbody>{talents}</tbody>
				</table>
			</div>
		);
	}
	return builds;
};

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
		if (!this.props.build) {
			if (search === '') {
				search = '?';
			} else {
				search += '&';
			}
			search += 'build=' + this.props.Builds[0].ID;
		} else if (search.indexOf('build') === -1) {
			return;
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
			winrates = <Winrates winrates={rates} />;
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
						<span>{wr.hero.Name}</span>
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

function pct(x) {
	return x.toFixed(1) + '%';
}

function toLength(l) {
	const mins = Math.trunc(l / 60);
	let secs = (l % 60).toString();
	if (secs.length < 2) {
		secs = '0' + secs;
	}
	return mins + ':' + secs;
}

// Modified from https://www.quirksmode.org/js/cookies.html.
function createCookie(name, value, days) {
	let expires = '';
	if (days) {
		const date = new Date();
		date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
		expires = '; expires=' + date.toGMTString();
	}
	document.cookie = name + '=' + value + expires + '; path=/';
}

function readCookie(name) {
	const nameEQ = name + '=';
	const ca = document.cookie.split(';');
	for (let i = 0; i < ca.length; i++) {
		let c = ca[i];
		while (c.charAt(0) === ' ') {
			c = c.substring(1, c.length);
		}
		if (c.indexOf(nameEQ) === 0) {
			return c.substring(nameEQ.length, c.length);
		}
	}
	return null;
}

const App = withRouter(HotsApp);

export default App;
