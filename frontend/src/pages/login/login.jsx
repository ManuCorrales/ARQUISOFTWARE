const Login = () => {

    return(
        <div style={{ 
            width:" 100% ", 
            height:" 100% ", 
            zIndex:5, 
            justifyContent: "center", 
            display:" flex " 
            }}>
            <div style={{ 
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
                    <input  style={{ right:" 10px ", position: " absolute"}}/> 
                </div>
                <div style={{ 
                    paddingBottom:" 7px ",
                    width:" 100% ",
                    position:" Relative "
                }}>  
                    <label>  Contraseña:  </label>
                   <input type="password" style={{ position:" absolute ", right:" 10px " }}/>
                </div>   
                <div style={{ width:" 100% ", margin:" 7px ", justifyContent:" center ", display:" Flex " }}>
                     <button  > Enviar </button>
                </div>
            </div> 
            {/* list/card component  (componente que hace la lista de hoteles) */}
        </div>
    );
}

export default Login;
