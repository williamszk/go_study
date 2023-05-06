import React, { Fragment, useEffect, useState } from 'react';
import './App.css';

function App(props) {
	const [isTrue, setIsTrue] = useState(true);
	const [crowd, setCrowd] = useState([]);

	const toggleTrue = () => {
		if (isTrue) {
			setIsTrue(false);
			return;
		}
		setIsTrue(true);
	};

	useEffect(() => {
		console.log('useEffect fired...');

		let people = [
			{
				id: 1,
				firstName: 'Mary',
				lastName: 'Jones',
				dob: '1964-10-10',
			},
			{
				id: 2,
				firstName: 'Bob',
				lastName: 'Ferguson',
				dob: '1964-10-12',
			},
		];
		// dob: data of birth

		setCrowd(people);
	}, []);

	return (
		<Fragment>
			<hr />
			<h1 className="h1-green">{props.msg}</h1>
			<hr />
			{isTrue && (
				<Fragment>
					<p>The current value of isTrue is true</p>
					<hr />
				</Fragment>
			)}
			<hr />

			{isTrue ? <p>Is true</p> : <p>Is false</p>}

			<hr />
			<a href="#!" className="btn btn-outline-secondary" onClick={toggleTrue}>
				Toggle isTrue
			</a>

			<hr />

			<h3>People</h3>

			<ul className="list-group">
				{crowd.map((m) => (
					<li key={m.id} className="list-group-item">
						{m.firstName} {m.lastName}
					</li>
				))}
				{/* {crowd.map((m) => {
					return (
						<li key={m.id} className="list-group-item">
							{m.firstName} {m.lastName}
						</li>
					);
				})} */}
			</ul>
		</Fragment>
	);
}

export default App;
