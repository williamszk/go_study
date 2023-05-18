import React from 'react';
import ReactDOM from 'react-dom/client';
// import App from './App';
import AppClass from './AppClass.js';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
	<React.StrictMode>
		<div className="container">
			<div className="row">
				<div className="col">
					<AppClass msg="Hello, World from a functional component" />
					{/* <App msg="Hello, World from a functional component" /> */}
				</div>
			</div>
		</div>
	</React.StrictMode>
);
