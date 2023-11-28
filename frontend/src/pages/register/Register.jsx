import axios from "axios";
import React, { useState } from "react";



const Register = () => {
    const [formValues, setFormValues] = useState({
        name: "",
        last_name: "",
        user_name: "",
        password: "",
        email: "",
        isAdmin: false
      });

      const handleChange = (event) => {
        const { name, value } = event.target;
        setFormValues((prevValues) => ({
          ...prevValues,
          [name]: value,
        }));
      };

      function submitForm(){
        console.log( formValues )
        axios.post("http://localHost:8090/user", formValues).then( function (res){console.log(res)}).catch(err=>{console.log(err)} )

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
                background:" #d5d8da ",
                borderColor:"black", 
                padding:" 15px ",
                border:" 1px solid black"
                }}>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>    
                    <label> Usuario: </label>
                    <input  
                            type="text"
                            name="user_name"
                            value={formValues.user_name} 
                            onChange={handleChange}
                            style={{ right:" 10px ", position: " absolute"}}/> 
                </div>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>    
                    <label> Nombre: </label>
                    <input  type="text"
                            name="name"
                            value={formValues.name}
                            style={{ right:" 10px ", position: " absolute"}}
                            onChange={handleChange}/> 
                </div>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>    
                    <label> Apellido: </label>
                    <input  
                            type="text"
                            name="last_name"
                            value={formValues.last_name}
                            style={{ right:" 10px ", position: " absolute"}}
                            onChange={handleChange}/> 
                </div>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>    
                    <label> Email: </label>
                    <input   type="text"
                             name="email"
                             value={formValues.email}
                             onChange={handleChange} 
                             style={{ right:" 10px ", 
                                      position: " absolute"}}/> 
                </div>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>  
                    <label>  Contraseña:  </label>
                   <input    type="password"
                             name="password"
                             value={formValues.password}
                             onChange={handleChange}
                            style={{ position:" absolute ", right:" 10px " }}/>
                </div>   
                <div style={{ width:" 100% ", margin:" 7px ", justifyContent:" center ", display:" Flex " }}>
                     <button  onClick={submitForm}> Registrar </button>
                </div>
            </div> 
            {/* list/card component  (componente que hace la lista de hoteles) */}
        </div>
    );
}

export default Register;
