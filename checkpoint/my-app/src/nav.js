import React from 'react';
import {Link} from 'react-router-dom';
import './nav.css';

function Nav(){
    return (
        <nav>
          <h3>Menu</h3>
            <ul className="nav_links">
                <Link className="nav_link" to="/Home">
                <li>Home</li>
                </Link>
                <Link className="nav_link" to="/Applications">
                <li>Applications</li>
                </Link>
            </ul>
        </nav>
    )
}

export default Nav;