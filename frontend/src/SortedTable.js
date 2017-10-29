import React, { Component } from 'react';
import { createCookie, readCookie } from './common';

const tableSortings = {};
const sortTableCookie = 'sort-table-';
function setTableSort(name, sort, dir) {
	const val = sort + ',' + (dir === true);
	createCookie(sortTableCookie + name, val);
	const sortings = tableSortings[name];
	const st = {
		sort: sort,
		sortDir: dir,
	};
	sortings.forEach(v => {
		v.setState(st);
	});
}
function registerTableSort(that) {
	const name = that.props.name;
	if (!tableSortings[name]) {
		tableSortings[name] = [];
	}
	tableSortings[name].push(that);
}
function unregisterTableSort(that) {
	const name = that.props.name;
	tableSortings[name] = tableSortings[name].filter(v => v !== that);
}

class SortedTable extends Component {
	constructor(props) {
		super(props);
		const lookup = this.lookup(props.headers);
		let sort = props.sort;
		if (!this.props.name || !lookup[props.sort]) {
			debugger;
		}
		let dir = lookup[props.sort].desc === true;
		const cookie = readCookie(sortTableCookie + this.props.name);
		registerTableSort(this);
		if (cookie) {
			const sp = cookie.split(',', 2);
			if (sp.length === 2) {
				sort = sp[0];
				dir = sp[1] === 'true';
			}
		}
		this.state = {
			lookup: lookup,
			sort: sort,
			sortDir: dir,
		};
		this.sort = this.sort.bind(this);
		this.sortClass = this.sortClass.bind(this);
	}
	componentWillUnmount() {
		unregisterTableSort(this);
	}
	lookup(headers) {
		const lookup = {};
		headers.forEach(h => {
			if (!h.cell) {
				h.cell = v => v;
			}
			if (!h.header) {
				h.header = h.name;
			}
			lookup[h.name] = h;
		});
		return lookup;
	}
	componentWillUpdate(nextProps, nextState) {
		nextState.lookup = this.lookup(nextProps.headers);
	}
	sort(sort) {
		if (this.state.lookup[sort].cmp === null) {
			return;
		}
		let dir;
		if (this.state.sort === sort) {
			dir = !this.state.sortDir;
		} else {
			dir = this.state.lookup[sort].desc === true;
		}
		setTableSort(this.props.name, sort, dir);
	}
	sortClass(col) {
		if (col !== this.state.sort) {
			return '';
		}
		const dir = this.state.sortDir ? 'desc' : 'asc';
		return 'sort-' + dir;
	}
	render() {
		const data = this.props.data;
		const cmp = this.state.lookup[this.state.sort].cmp || defaultCmp;
		data.sort((a, b) => cmp(a[this.state.sort], b[this.state.sort]));
		if (this.state.sortDir) {
			data.reverse();
		}
		const body = data.map((row, i) => (
			<tr key={i}>
				{this.props.headers.map(h => (
					<td key={h.name}>{h.cell(row[h.name])}</td>
				))}
			</tr>
		));
		return (
			<table className="sorted">
				<thead>
					<tr>
						{this.props.headers.map(h => {
							let { name, title, header } = h;
							return (
								<th
									key={name}
									onClick={() => this.sort(name)}
									className={this.sortClass(name)}
									title={title}
								>
									{header}
								</th>
							);
						})}
					</tr>
				</thead>
				<tbody>{body}</tbody>
			</table>
		);
	}
}

function defaultCmp(a, b) {
	switch (typeof a) {
		case 'number':
		case 'boolean':
			return a - b;
		case 'string':
			return a.localeCompare(b);
		default:
			if (a instanceof Date) {
				return a - b;
			} else {
				debugger;
			}
	}
}

export default SortedTable;
