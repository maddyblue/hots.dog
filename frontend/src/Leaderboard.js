// @flow

import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import {
	Fetch,
	regionNames,
	readCookie,
	regionCookie,
	GameModes,
	Title,
} from './common';
import SortedTable from './SortedTable';

type Props = {
	handleChange: any,
	region: string,
	mode: string,
	build?: string,
	Builds: any,
	Modes: any,
};

class Leaderboard extends Component<
	Props,
	{
		Players: ?(any[]),
		search: string,
		Days: number,
		MinGames: number,
	}
> {
	state = {
		Players: null,
		search: '',
		MinGames: 0,
		Days: 0,
	};
	constructor(props: Props) {
		super(props);
		if (!props.region) {
			this.props.handleChange(
				{
					target: { name: 'region', value: readCookie(regionCookie) || '1' },
				},
				true
			);
		}
		if (!props.mode) {
			this.props.handleChange(
				{
					target: { name: 'mode', value: '3' },
				},
				true
			);
		}
	}
	componentDidMount() {
		this.update();
	}
	componentDidUpdate() {
		this.update();
	}
	makeSearch = () => {
		if (!this.props.mode || !this.props.region) {
			return;
		}
		return (
			'/api/get-leaderboard?mode=' +
			this.props.mode +
			'&region=' +
			this.props.region
		);
	};
	update = () => {
		const search = this.makeSearch();
		if (!search) {
			return;
		}
		if (this.state.search === search) {
			return;
		}
		this.setState({ Players: null, search: search });
		Fetch(search, data => {
			if (search === this.makeSearch()) {
				if (!data.Players) {
					data.Players = [];
				}
				this.setState(data);
			}
		});
	};
	render() {
		let content;
		if (this.state.Players === null) {
			content = 'loading...';
		} else if (this.state.Players && !this.state.Players.length) {
			content = 'No leaderboard data.';
		} else {
			content = (
				<div>
					<p>
						Top players in Master League who have played at least{' '}
						{this.state.MinGames} games in the last {this.state.Days} days.
					</p>
					<SortedTable
						name="leaderboard"
						sort="Rank"
						headers={[
							{
								name: 'Rank',
							},
							{
								name: 'Battletag',
								cell: (v, row) => (
									<Link
										to={'/players/' + this.props.region + '/' + row.Blizzid}
									>
										{v}
									</Link>
								),
							},
							{
								name: 'Skill',
								cell: v => v.toFixed(6),
							},
							{
								name: 'Games',
							},
						]}
						data={this.state.Players || []}
					/>
				</div>
			);
		}
		return (
			<div>
				<Title {...this.props} prefix="Leaderboard" />
				<div className="row">
					<div className="column">
						<label>Region</label>
						<select
							name="region"
							value={this.props.region}
							onChange={this.props.handleChange}
						>
							{Object.keys(regionNames)
								.sort()
								.map(k => (
									<option key={k} value={k}>
										{regionNames[k]}
									</option>
								))}
						</select>
					</div>
					<div className="column">
						<label>Game Mode</label>
						<select
							name="mode"
							value={this.props.mode}
							onChange={this.props.handleChange}
						>
							<GameModes Modes={this.props.Modes} noAll={true} />
						</select>
					</div>
				</div>
				{content}
			</div>
		);
	}
}

export { Leaderboard };
