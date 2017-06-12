import React, { Component } from 'react';
//import logo from './logo.svg';
import './App.css';
import './tachyons.min.css';

class App extends Component {
	constructor(props) {
		super(props);
		var state = {
			Builds: [],
			Heroes: [],
			Maps: [],
			Modes: {},
			Winrates: {},
		};
		this.params = [
			'build',
			'map',
			'mode',
		];
		var search = new URLSearchParams(window.location.search);
		this.params.forEach(key => {
			if (!search.has(key)) {
				return;
			}
			state[key] = search.get(key);
		});
		this.state = state;
		this.handleChange = this.handleChange.bind(this);
	}
	componentDidMount() {
		fetch('/api/init').then(resp => resp.json().then(data => {
			if (!this.state.build) {
				data.build = data.Builds[0].ID;
			}
			this.setState(data, this.update);
		}));
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
		return params.join('&');
	}
	update() {
		var search = this.getSearch();
		fetch('/api/get-winrates?' + search).then(resp => resp.json().then(data => {
			if (this.getSearch() === search) {
				this.setState({Winrates: data});
			}
		}));
	}
	handleChange(event) {
		var st = {};
		st[event.target.name] = event.target.value;
		this.setState(st, () => {
			var params = this.getSearch();
			window.history.replaceState(null, null, '/?' + params);
			this.update();
		});
	}
	render() {
		if (!this.state.Builds.length) {
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
				<select name="map" value={this.state.map} onChange={this.handleChange}>
					{maps}
				</select>
				<select name="build" value={this.state.build} onChange={this.handleChange}>
					{builds}
				</select>
				<select name="mode" value={this.state.mode} onChange={this.handleChange}>
					{modes}
				</select>
				<div>
					<Winrates winrates={this.state.Winrates} heroes={this.state.Heroes} />
				</div>
			</div>
		);
	}
}

const Winrates = (props) => {
	if (!props.winrates.Current) {
		return null;
	}
	var winrates = props.heroes.map(hero => {
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

export default App;
