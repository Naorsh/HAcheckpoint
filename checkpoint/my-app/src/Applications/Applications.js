import React,{useEffect, useState} from 'react';
import {useSelector} from 'react-redux';
import App from './../app/app';
import AppsTableHeader from './appsTableHeader/AppsTableHeader';
import AddAppForm from './addAppForm/AddAppForm';
import './Applications.css';


function Applications() {
  const formSubmitted = useSelector(state => state.form); 
  const [applications, setApplications] = useState([]);
  
  useEffect(() => {
    getApplications();
  },[formSubmitted]);

  const getApplications = async () => {
    await fetch('/live');
    const response = await fetch('/api/getApplications',{
      method: 'GET',
    });
    const data = await response.json();
    setApplications(data);
  }

  return (
    <div className="App">
      <div className="applications">
        <div className="application-header">
          <h1 className="main-header">Applications List</h1>
        </div>
        <div className="add_app_form">
          <AddAppForm />
        </div>
        <div className="applications-list">
          <div className="app-table-header">
            <AppsTableHeader />
          </div>
          {applications &&
            <div className="app-table-body">
            {applications.map(app =>(
              <App 
              key={app.ID} 
              id={app.ID} 
              name={app.Name}
              appKey={app.Key}
              creationTime={app.CreationTime}/>
            ))}
          </div>
          }
          
        </div>
      </div>
    </div>
  );
}

export default Applications;
