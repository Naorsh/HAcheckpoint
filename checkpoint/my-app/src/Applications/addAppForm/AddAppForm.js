import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { formSubmittedSuccesfully } from './../../actions'
import './AddAppForm.css'

const AddAppForm = () => {
    const [name, setName] = useState("");
    const [key, setKey] = useState("");
    const dispatch = useDispatch();

    const postApplication = async () => {
        try {
            await fetch("/api/addApplication", {
                method: 'post',
                mode: 'cors',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    'name': name,
                    'key': key
                })
            });
            dispatch(formSubmittedSuccesfully())
            setName("");
            setKey("");
        } catch (e) {
            console.log(e)
        } 
    }

    const addAppToDB = e => {
        e.preventDefault();
        postApplication();

    }


    const handleNameChange = e => {
        setName(e.target.value)
    }

    const handleKeyChange = e => {
        setKey(e.target.value)
    }

    return (
        <div className="application_form">
            <div className="add_application_text" >Add Application:</div>
            <form onSubmit={addAppToDB} className="add-app-form">
                <input className="input_textbox" value={name} onChange={handleNameChange} maxLength="15" required placeholder="Application Name" />
                <input className="input_textbox" value={key} onChange={handleKeyChange} maxLength="15" required placeholder="Application Key" />
                <button className="add-app-submit-button" type="submit">
                    Add Application</button>
            </form>
        </div>
    )
}

export default AddAppForm;