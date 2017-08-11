import React, { Component } from 'react';
import {
	Route,
//	Link,
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
		maps.unshift(<option key=""></option>);
		var builds = this.state.Builds.map(b => <option key={b.ID} value={b.ID}>
			{b.Patch} ({new Date(b.Start).toLocaleDateString()} - {new Date(b.Finish).toLocaleDateString()})
		</option>);
		var modeKeys = Object.keys(this.state.Modes);
		modeKeys.sort().reverse();
		var modes = modeKeys.map(k => <option key={k} value={k}>{this.state.Modes[k]}</option>);
		modes.unshift(<option key=""></option>);
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
				</div>
		);
	}
}

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
		var search = window.location.search;
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
		return <Winrates winrates={this.state.winrates} Heroes={this.props.Heroes}/>;
	}
}

const Winrates = (props) => {
	var winrates = props.Heroes.map(hero => {
		var wr = props.winrates.Current[hero];
		var games, rate, prevRate;
		if (wr) {
			games = wr.Wins + wr.Losses;
			var winRate = (wr.Wins / games * 100);
			rate = winRate.toFixed(1) + '%';
			var prev = props.winrates.Previous[hero];
			if (prev) {
				var prevGames = prev.Wins + prev.Losses;
				var prevWinRate = prev.Wins / prevGames * 100;
				prevRate = (winRate - prevWinRate).toFixed(1) + '%';
			}
		}
		return (
			<tr key={hero}>
				<td className="pv1">{hero}</td>
				<td>{games}</td>
				<td>{rate}</td>
				<td>{prevRate}</td>
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
					<th className="pv2 ph3">change since previous patch</th>
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
