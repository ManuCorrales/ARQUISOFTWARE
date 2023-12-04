import axios from "axios";
import React, { useState } from "react";
import { useEffect } from "react";



function ListaReservas(){
    //const [formValues, setFormValues] = useState(fetch('http://localhost:8090/hotels'))
    const [hotels, setHotels] = useState([])
    const[hotelSeleccionado, setHotelSeleccionado] = useState({})
    const[reservas, setReservas] = useState([])
    const [fechaSeleccionada, setFechaSeleccionada] = useState("");
    const handleHotelChange = (event) => {
        const selectedValue = event.target.value;
        setHotelSeleccionado(selectedValue);
        console.log(selectedValue)
      };
    const handleFechaChange = (event) => {
        const selectedDate = event.target.value;
        setFechaSeleccionada(selectedDate);
      };

    useEffect(() => {
      fetch('http://localhost:8090/hotels')
        .then((response) => {
          return response.json()
        })
        .then((dataRequest) => {
          setHotels(dataRequest)
        })
    }, [])

    useEffect(()=>{
        console.log(hotels)
    }, [hotels])

    useEffect(() => {
        fetch('http://localhost:8090/reservas')
          .then((response) => {
            return response.json()
          })
          .then((dataRequest) => {
            setReservas(dataRequest)
            console.log(dataRequest)
          })
      }, [])


      // Filtrar las reservas por hotelSeleccionado
  // Filtrar las reservas por hotel y fecha seleccionada
const reservasFiltradas = reservas.filter((reserva) => {
    const matchHotel = hotelSeleccionado === "todos" || reserva.hotel_id == hotelSeleccionado;
    const matchFecha = fechaSeleccionada === "" || (
      new Date(reserva.date_from) <= new Date(fechaSeleccionada) &&
      new Date(fechaSeleccionada) <= new Date(reserva.date_to)
    );
    return matchHotel && matchFecha;
  });
  
  
  return(
  <div style={{justifyContent:"center", display:"flex", flexDirection:"column"}} >
    <div>
  <label for="hotel">Selecciona un hotel:</label>

  <select name="hotel" id="hotel" onChange={handleHotelChange}>     
    <option key={-1} value={"todos"}> Todos </option> 
    {hotels.map((item, index) => (
          <option key={index} value={item.id}>
            {item.name}
          </option>
        ))}
  </select>
  <label htmlFor="fecha">Selecciona una fecha:</label>
      <input type="date" id="fecha" name="fecha" onChange={handleFechaChange} value={fechaSeleccionada} />
  </div>
  <div>
  <table>
        <thead>
          <tr>
            <th>Date From</th>
            <th>Date To</th>
            <th>Hotel ID</th>
          </tr>
        </thead>
        <tbody>
          {reservasFiltradas.map((reserva, index) => (
            <tr key={index}>
              <td>{reserva.date_from}</td>
              <td>{reserva.date_to}</td>
              <td>{reserva.hotel_id}</td>
            </tr>
          ))}
        </tbody>
      </table>
      </div>
  </div>
  )
}

export default ListaReservas;


