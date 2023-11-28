import axios from "axios";
import {useState} from 'react';
import { useNavigate } from 'react-router-dom';


function Login(){

const [message, setMessage] = useState('');
const [formValues, setFormValues] = useState({
    user_name: "",
    password: "",
  });
  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormValues((prevValues) => ({
      ...prevValues,
      [name]: value,
    }));
  };

function submitLogin(){
     console.log(formValues)
axios.post('http://localHost:8090/login', formValues)
     .then(function (res){localStorage.setItem("userData", JSON.stringify(res.data));window.location.href = '/';}).catch(err=>{console.log(err)})
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
                background: "#d5d8da",
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
                    value={formValues.user_name}
                    name={"user_name"}
                    /> 
                    
                </div>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>  
                    <label>  Contraseña:  </label>
                   <input type="password"
                   onChange={handleChange}
                        style={{ position:" absolute ", right:" 10px " }}
                        name={"password"}/>
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
