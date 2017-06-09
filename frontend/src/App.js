import React, { Component } from 'react';
//import logo from './logo.svg';
import './App.css';
import './tachyons.min.css';

class App extends Component {
	constructor(props) {
		super(props);
		this.state = {
			Builds: [],
			Heroes: [],
			Maps: [],
			Modes: {},
			Winrates: {},
		};
		this.handleChange = this.handleChange.bind(this);
	}
	componentDidMount() {
		fetch('/api/init').then(resp => resp.json().then(data => {
			this.setState(data, this.update);
		}));
	}
	update(state) {
		var url = '/api/get-winrates?';
		[
			'build',
			'map',
			'mode',
		].forEach(key => {
			var value = this.state[key];
			if (!value) {
				return;
			}
			url += '&' + key + '=' + encodeURIComponent(value);
		});
		fetch(url).then(resp => resp.json().then(data => {
			this.setState({Winrates: data});
		}));
	}
	handleChange(event) {
		var st = {};
		st[event.target.name] = event.target.value;
		this.setState(st, this.update);
	}
	render() {
		var maps = this.state.Maps.map(m => <option key={m}>{m}</option>);
		maps.unshift(<option key=""></option>);
		var builds = this.state.Builds.map(b => <option key={b.ID} value={b.ID}>
			{b.Patch} ({new Date(b.Start).toLocaleDateString()} - {new Date(b.Finish).toLocaleDateString()})
		</option>);
		var modeKeys = Object.keys(this.state.Modes);
		modeKeys.sort().reverse();
		var modes = modeKeys.map(k => <option key={k} value={k}>{this.state.Modes[k]}</option>);
		modes.unshift(<option key=""></option>);
		var winrates = Object.keys(this.state.Winrates).map(hero => {
			var wr = this.state.Winrates[hero];
			return (
				<tr key={hero}>
					<td className="pv1">{hero}</td>
					<td>{wr.Wins}</td>
					<td>{wr.Losses}</td>
				</tr>
			);
		});
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
					<table className="ba br2 b--black-10 pv2 ph3">
						<thead>
							<tr>
								<th className="pv2 ph3">hero</th>
								<th className="pv2 ph3">wins</th>
								<th className="pv2 ph3">losses</th>
							</tr>
						</thead>
						<tbody>
							{winrates}
						</tbody>
					</table>
				</div>
			</div>
		);
	}
}

export default App;
