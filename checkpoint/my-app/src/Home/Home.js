import React from 'react';
import './Home.css';
import {useSelector, useDispatch} from 'react-redux';
import {changeColorToRed, changeColorToBlue, changeColorToBlack} from './../actions'

function Home() {

    const color = useSelector(state => state.color);
    const dispatch = useDispatch();

    function onColorSelect(event){
        switch(event.target.value) {
            case 'red':
                dispatch(changeColorToRed())
                break;
            case 'blue':
                dispatch(changeColorToBlue())
                break;
            case 'black':
                dispatch(changeColorToBlack())
                break;
            default:
                break;
          }
    }
    
  return (
    <div className="Home">
        <div className="header">
            <h1 className={color}>Hello Guest</h1>
        </div>
        <div className="colors-ddl">
            <select className={color} name="colors" id="colors" onChange={onColorSelect} defaultValue={color ? color : "black"}>
                <option value="black" className="black">Black</option>
                <option value="red" className="red">Red</option>
                <option value="blue" className="blue">Blue</option>
            </select>
        </div>
    </div>
  );
}

export default Home;