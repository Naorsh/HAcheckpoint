import React from 'react';
import './AppsTableHeader.css'

const AppsTableHeader = () => {
    return(
        <div className="header_row">
            <div className="id_col">ID</div>
            <div className="name_col">Name</div>
            <div className="key_col">Key</div>
            <div className="ct_col">Creation Time</div>
        </div>
    )
}

export default AppsTableHeader;