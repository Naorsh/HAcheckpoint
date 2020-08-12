import React from 'react';
import './app.css'

const App = ({id, name, appKey, creationTime}) => {
    return(
        <div className="app_row">
            <div className="app_id">{id}</div>
            <div className="app_name">{name}</div>
            <div className="app_key">{appKey}</div>
            <div className="app_ct">{creationTime}</div>
        </div>
    )
}

export default App;