import React, { Fragment, useEffect, useRef, useState } from 'react';
import './App.css';
import Input from './Input';

function App(props) {
	const [isTrue, setIsTrue] = useState(true);
	const [crowd, setCrowd] = useState([]);

	const [firstName, setFirstName] = useState('');
	const [lastName, setLastName] = useState('');
	const [dateOfBirth, setDateOfBirth] = useState('');

	// refs
	const firstNameRef = useRef();
	const lastNameRef = useRef(null);
	const dateOfBirthRef = useRef(null);

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

	const handleSubmit = (event) => {
		// this will stop the default action of submit
		// which is to make an http request to the server
		event.preventDefault();

		// console.log(firstName, lastName, dateOfBirth);
		if (lastName !== '') {
			addPerson(firstName, lastName, dateOfBirth);
		}
	};

	const addPerson = (newFirst, newLast, newDateOfBirth) => {
		let newPerson = {
			id: this.crowd.length + 1,
			firstName: newFirst,
			lastName: newLast,
			dateOfBirth: newDateOfBirth,
		};

		const newList = this.crowd.concat(newPerson);

		const sorted = newList.sort((a, b) => {
			if (a.lastName < b.lastName) {
				return -1;
			} else if (a.lastName > b.lastName) {
				return 1;
			}
			return 0;
		});

		setCrowd(sorted);
		// reset all global variables
		setFirstName('');
		setLastName('');
		setDateOfBirth('');

		// console.log('crowd:', crowd);

		firstNameRef.current.value = '';
		lastNameRef.current.value = '';
		dateOfBirthRef.current.value = '';
	};

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

			<form autoComplete="off" onSubmit={handleSubmit}>
				<div className="mb-3">
					<label className="form-label" htmlFor="first-name">
						First name
					</label>

					<input
						type="text"
						name="first-name"
						id="first-name"
						ref={firstNameRef}
						autoComplete="first-name-new"
						className="form-control"
						onChange={(event) => {
							setFirstName(event.target.value);
						}}
					></input>
				</div>

				<Input
					title="Last Name"
					type="text"
					ref={lastNameRef}
					name="last-name"
					autoComplete="last-name-new"
					className="form-control"
					onChange={(event) => setLastName(event.target.value)}
				></Input>

				<Input
					title="Date of Birth"
					type="date"
					ref={dateOfBirthRef}
					name="date-of-birth"
					autoComplete="date-of-birth-new"
					className="form-control"
					onChange={(event) => setDateOfBirth(event.target.value)}
				></Input>

				<input type="submit" value="Submit" className="btn btn-primary"></input>
			</form>

			<div>
				First Name: {firstName} <br />
				Last Name: {lastName} <br />
				Date of Birth: {dateOfBirth} <br />
			</div>

			<hr />
			<h3>People</h3>

			<ul className="list-group">
				{this.crowd.map((m) => (
					<li key={m.id} className="list-group-item">
						{m.firstName} {m.lastName}
					</li>
				))}
			</ul>
		</Fragment>
	);
}

export default App;
