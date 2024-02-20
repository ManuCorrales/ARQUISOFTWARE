import axios from "axios";
import React, { useState } from "react";
import { useEffect } from "react";



function ListaReservas(){
    //const [formValues, setFormValues] = useState(fetch('http://localhost:8090/hotels'))
    const [hotels, setHotels] = useState([])
    const [usuarios, setUsuarios] = useState([])
    const[hotelSeleccionado, setHotelSeleccionado] = useState("todos")
    const[reservas, setReservas] = useState([])
    const [fechaSalida, setFechaSalida] = useState("");
    const [fechaLlegada, setFechaLlegada] = useState("");
    const [reservasFiltradas, setReservasFiltradas] = useState([])


    const handleDateChange = (event) => {

    }

    const handleHotelChange = (event) => {
        const selectedValue = event.target.value;
        setHotelSeleccionado(selectedValue);
      };
    const handleFechaSalidaChange = (event) => {
        const selectedDate = event.target.value;
        setFechaSalida(selectedDate);

        let filtradas = []
        reservas.forEach(reserva => {
          console.log(hotelSeleccionado);
          if(hotelSeleccionado === "todos" || reserva.hotel_id == hotelSeleccionado){
            if(fechaLlegada === "" && fechaSalida === "" ||  (new Date(reserva.date_from) <= new Date(fechaSalida) && new Date(fechaLlegada) <= new Date(reserva.date_to)))
              filtradas.push(reserva);
          }
        });
        
        setReservasFiltradas(filtradas);
        console.log(reservasFiltradas);
      };
      
    const handleFechaLlegadaChange = (event) => {
      const selectedDate = event.target.value;
      setFechaLlegada(selectedDate);

      let filtradas = []
      reservas.forEach(reserva => {
        if(hotelSeleccionado === "todos" || reserva.hotel_id == hotelSeleccionado){
          if(fechaLlegada === "" ||  (new Date(reserva.date_from) >= new Date(fechaLlegada) && new Date(fechaSalida) >= new Date(reserva.date_to)))
            filtradas.push(reserva);
        }
      });
      
      setReservasFiltradas(filtradas);
  
    };
      const [userData, setUserData] = useState(() => {
        // getting stored value
        const saved = localStorage.getItem("userData");
        const initialValue = JSON.parse(saved);
        return initialValue || "";
      });

    useEffect(() => {
      fetch('http://localhost:8090/hotels')
        .then((response) => {
          return response.json()
        })
        .then((dataRequest) => {
          setHotels(dataRequest)
        })
    }, [])
    useEffect(() => {
        fetch('http://localhost:8090/users')
          .then((response) => {
            return response.json()
          })
          .then((dataRequest) => {
            setUsuarios(dataRequest)
            console.log(dataRequest)
          })
      }, [])

    useEffect(() => {
        fetch('http://localhost:8090/reservas')
          .then((response) => {
            return response.json()
          })
          .then((dataRequest) => {
            setReservas(dataRequest)
          })
      }, [])


      // Filtrar las reservas por hotelSeleccionado
  // Filtrar las reservas por hotel y fecha seleccionada


  
  
  return(
  <div style={{
    backgroundColor:"white",
    borderRadius:"10px",
    justifyContent:"center", 
    display:"flex", 
    flexDirection:"column", 
    marginTop:"30px", 
    margin:"50px"
    }} >
  <div style={{
    margin:"20px",
    display:"grid",
    gridTemplateColumns:"1fr, 1fr"
  }}>
  <div style={{
    gridColumn:"1",
    marginRight:"20px",
    borderColor:"black",
    borderRadius:"10px",
    borderStyle:"solid",
    borderWidth:"1px",
    padding:"10px"
  }}>
      <label htmlFor="hotel">Selecciona un hotel:</label>
  <select name="hotel" id="hotel" onChange={handleHotelChange}>     
    <option key={-1} value={"todos"}> Todos </option> 
    {hotels.map((item, index) => (
          <option key={index} value={item.id}>
            {item.name}
          </option>
        ))}
  </select>
  </div>
  <div style={{
    gridColumn:"2",
    marginRight:"20px",
    borderColor:"black",
    borderRadius:"10px",
    borderStyle:"solid",
    borderWidth:"1px",
    padding:"10px",
    justifyContent:"center"
  }}>
  <label htmlFor="fecha_llegada">Selecciona una fecha de llegada:</label>
      <input type="date" id="fecha_llegada" name="fecha_llegada" onChange={handleFechaLlegadaChange} value={fechaLlegada} />
      
  <label htmlFor="fecha_salida">Selecciona una fecha de salida:</label>
      <input type="date" id="fecha_salida" name="fecha_salida" onChange={handleFechaSalidaChange} value={fechaSalida} />
 </div> 
 </div>
  <div style={{display:"flex", flexDirection:"column"}}>
  <table style={{
    marginLeft:"80px",
    marginRight:"80px",
    marginBottom:"80px"
  }}>
        <thead>
          <tr>
            <th style={{textAlign:"left",
                        marginRight:"20px",
                        borderColor:"black",
                        borderStyle:"solid",
                        borderWidth:"1px",
                        }}>Date From</th>
            <th style={{textAlign:"left",
                        marginRight:"20px",
                        borderColor:"black",
                        borderStyle:"solid",
                        borderWidth:"1px",
                         }}>Date To</th>
            <th style={{textAlign:"left",
                        marginRight:"20px",
                        borderColor:"black",
                        borderStyle:"solid",
                        borderWidth:"1px",
                        }}>Hotel</th>
            {userData.isadmin&&<th style={{textAlign:"left",
                        marginRight:"20px",
                        borderColor:"black",
                        borderStyle:"solid",
                        borderWidth:"1px",
                        }}>Usuario & Email</th>}
          </tr>
        </thead>
        <tbody>
          {reservasFiltradas.map((reserva, index) => (
            <tr key={index}>
              <td style={{textAlign:"left",
                        marginRight:"20px",
                        borderColor:"black",
                        borderStyle:"solid",
                        borderWidth:"1px",
                        }}>{reserva.date_from}</td>
              <td style={{textAlign:"left",
                        marginRight:"20px",
                        borderColor:"black",
                        borderStyle:"solid",
                        borderWidth:"1px",
                        }}>{reserva.date_to}</td>
              <td style={{textAlign:"left",
                        marginRight:"20px",
                        borderColor:"black",
                        borderStyle:"solid",
                        borderWidth:"1px",
                        }}>{hotels.filter((hotelj)=> hotelj.hotel_id==reserva.hotel_id).map((hoteli)=> (hoteli.name))}</td>
              {userData.isadmin&&<td style={{textAlign:"left",
                        marginRight:"20px",
                        borderColor:"black",
                        borderStyle:"solid",
                        borderWidth:"1px",
                        }}>{usuarios.filter((userj)=> userj.id==reserva.user_id).map((useri)=> (useri.user_name + " | " + useri.email))}</td>}
            </tr>
          ))}
        </tbody>
      </table>
      </div>
  </div>
  )
}

export default ListaReservas;


