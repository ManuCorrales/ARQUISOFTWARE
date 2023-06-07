import axios from "axios";
import {useState} from 'react';

function Login(){

const [message, setMessage] = useState('');


const handleChange = event => {
    setMessage(event.target.value);
  };

function submitLogin(){
axios.get('/user/username/'+message)
     .then(function (res){console.log(res)}).catch(err=>{console.log(err)})
}
    return(
        <div style={{ 
            width:" 100% ", 
            height:" 100% ", 
            zIndex:5, 
            justifyContent: "center", 
            display:" flex " 
            }}>
            <div style={{ 
                borderRadius:" 5px ", 
                width:" 350px ", 
                position: "relative",
                background:" #FACFCB ",
                borderColor:"black", 
                padding:" 15px ",
                border:" 1px solid black"
                }}>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>    
                    <label> User: </label>
                    <input  style={{ right:" 10px ", position: " absolute"}}
                    onChange={handleChange}
                    value={message}/> 
                </div>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>  
                    <label>  Contraseña:  </label>
                   <input type="password"
                        style={{ position:" absolute ", right:" 10px " }}/>
                </div>   
                <div style={{ width:" 100% ", margin:" 7px ", justifyContent:" center ", display:" Flex " }}>
                     <button onClick={submitLogin} > Enviar </button>
                </div>
            </div> 
            {/* list/card component  (componente que hace la lista de hoteles) */}
        </div>
    );
}

export default Login;
