import React, { Component } from 'react';
import {
	Route,
	Link,
	withRouter
} from 'react-router-dom';
//import logo from './logo.svg';
import './App.css';
import './tachyons.min.css';

class HotsApp extends Component {
	constructor(props) {
		super(props);
		this.params = [
			'build',
			'map',
			'mode',
		];
		this.state = this.searchState();
		this.handleChange = this.handleChange.bind(this);
	}
	componentDidMount() {
		fetch('/api/init').then(resp => resp.json().then(data => {
			data.init = true;
			if (!this.state.build) {
				data.build = data.Builds[0].ID;
				var search = window.location.search || '?';
				this.props.history.replace({search: search + 'build=' + data.build});
			}
			this.setState(data);
		}));
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
		var st = {};
		var search = new URLSearchParams(window.location.search);
		this.params.forEach(key => {
			if (!search.has(key)) {
				return;
			}
			st[key] = search.get(key);
		});
		return st;
	}
	getSearch() {
		var params = [];
		this.params.forEach(key => {
			var value = this.state[key];
			if (!value) {
				return;
			}
			value = encodeURIComponent(value)
			params.push(key + '=' + value);
		});
		return '?' + params.join('&');
	}
	handleChange(event) {
		var st = {};
		st[event.target.name] = event.target.value;
		this.setState(st, () => {
			var params = this.getSearch();
			this.props.history.push({search: params});
		});
	}
	render() {
		if (!this.state.build || !this.state.init) {
			return <div>loading...</div>;
		}
		var maps = this.state.Maps.map(m => <option key={m}>{m}</option>);
		maps.unshift(<option key="">All Maps</option>);
		var builds = this.state.Builds.map(b => <option key={b.ID} value={b.ID}>
			{b.Patch} ({new Date(b.Start).toLocaleDateString()} - {new Date(b.Finish).toLocaleDateString()})
		</option>);
		var modeKeys = Object.keys(this.state.Modes);
		modeKeys.sort().reverse();
		var modes = modeKeys.map(k => <option key={k} value={k}>{this.state.Modes[k]}</option>);
		modes.unshift(<option key="">All Game Modes</option>);
		return (
				<div className="sans-serif">
					<a href="/">home</a>
					<hr/>
					<select name="map" value={this.state.map} onChange={this.handleChange}>
						{maps}
					</select>
					<select name="build" value={this.state.build} onChange={this.handleChange}>
						{builds}
					</select>
					<select name="mode" value={this.state.mode} onChange={this.handleChange}>
						{modes}
					</select>

					<hr/>

					<Route exact path="/" render={props => <HeroWinrates params={this.params} {...this.state} {...props} />}/>
					<Route path ="/talents/:hero" render={props => <TalentWinrates params={this.params} {...this.state} {...props} />}/>
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
		fetch('/api/get-build-winrates/' + this.props.match.params.hero + search).then(resp => resp.json().then(data => {
			if (search === window.location.search) {
				this.setState({
					winrates: data,
					search: search,
				});
			}
		}));
	}
	render() {
		if (!this.state.winrates) {
			return <div>loading...</div>;
		}
		return <Builds winrates={this.state.winrates}/>;
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

const Builds = (props) => {
	var builds = [];
	for (let tier = 1; tier <= 7; tier++) {
		const curTier = props.winrates.Current[tier];
		const prevTier = props.winrates.Previous[tier];
		const talents = Object.keys(curTier).sort().map(talent => {
			const cur = curTier[talent];
			const prev = prevTier[talent];
			const games = cur.Wins + cur.Losses;
			const winRate = (cur.Wins / games * 100);
			const rate = winRate.toFixed(1) + '%';
			let change;
			if (prev) {
				const prevGames = prev.Wins + prev.Losses;
				const prevWinRate = prev.Wins / prevGames * 100;
				change = (winRate - prevWinRate).toFixed(1) + '%';
			}
			return (
				<tr key={talent}>
					<td className="pv1">{talent}</td>
					<td>{games}</td>
					<td>{rate}</td>
					<td>{change}</td>
				</tr>
			);
		});
		builds.push(
			<div key={tier}>
			Tier {tierNames[tier]}
			<table className="ba br2 b--black-10 pv2 ph3">
				<thead>
					<tr>
						<th className="pv2 ph3">hero</th>
						<th className="pv2 ph3">games</th>
						<th className="pv2 ph3">winrate</th>
						<th className="pv2 ph3" title="change since previous patch">change</th>
					</tr>
				</thead>
				<tbody>
					{talents}
				</tbody>
			</table>
			</div>
		);
	}
	return (
		<div>
			{builds}
		</div>
	);
};

class HeroWinrates extends Component {
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
		fetch('/api/get-winrates' + search).then(resp => resp.json().then(data => {
			if (search === window.location.search) {
				this.setState({
					winrates: data,
					search: search,
				});
			}
		}));
	}
	render() {
		if (!this.state.winrates) {
			return <div>loading...</div>;
		}
		return <Winrates winrates={this.state.winrates} search={this.state.search} Heroes={this.props.Heroes}/>;
	}
}

const Winrates = (props) => {
	var winrates = props.Heroes.map(hero => {
		const wr = props.winrates.Current[hero];
		let games, rate, change;
		if (wr) {
			games = wr.Wins + wr.Losses;
			var winRate = (wr.Wins / games * 100);
			rate = winRate.toFixed(1) + '%';
			const prev = props.winrates.Previous[hero];
			if (prev) {
				const prevGames = prev.Wins + prev.Losses;
				const prevWinRate = prev.Wins / prevGames * 100;
				change = (winRate - prevWinRate).toFixed(1) + '%';
			}
		}
		return (
			<tr key={hero}>
				<td className="pv1"><Link to={"/talents/" + hero + props.search}>{hero}</Link></td>
				<td>{games}</td>
				<td>{rate}</td>
				<td>{change}</td>
			</tr>
		);
	});
	return (
		<table className="ba br2 b--black-10 pv2 ph3">
			<thead>
				<tr>
					<th className="pv2 ph3">hero</th>
					<th className="pv2 ph3">games</th>
					<th className="pv2 ph3">winrate</th>
					<th className="pv2 ph3" title="change since previous patch">change</th>
				</tr>
			</thead>
			<tbody>
				{winrates}
			</tbody>
		</table>
	);
};

const App = withRouter(HotsApp);

export default App;
