import React from "react";
import "./AppFooter.css";

function AppFooter() {
	const currentYear = new Date().getFullYear();

	// .jsx uses xml behind, and it needs to close
	// this is why we need to write <hr/> instead of just <hr>
	return (
		<>
			<hr />
			<p className="footer">Copyright &copy; 1990 - {currentYear} Mario&Luigi Software Ltd.</p>
		</>
	);
}

export default AppFooter;
