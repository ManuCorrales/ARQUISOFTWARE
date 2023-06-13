import { useEffect, useState } from 'react'; 
import {config} from '../../config'

import Calendar from "@demark-pro/react-booking-calendar";



function Hotel() {
    const [lastDirectory, setLastDirectory] = useState('');
    const [hotel, setHotel] = useState({});

    const [selectedDates, setSelectedDates] = useState([]);//mandar al apretar reservar
    const handleChange = (e) => {setSelectedDates(e);}
    const[reserved, setReserved] = useState([
        {
          startDate: new Date(2023, 7, 10),
          endDate: new Date(2023, 7, 15),
        },
      ])

// desde el back, guardar y devolver las fechas reservadas asi:
//un objeto asi se manda para reservarlo, las reservas para bloquearlas se tienen que mandar en un array de estos. Al guardar una reserva nueva, añadirla al fin del array


  function reserveDate(){
    var temp=reserved
    temp.push({startDate:selectedDates[0],endDate:selectedDates[1]}) //en vez de guardar esto localmente se manda al back y ahi se añade al array de reservas
    console.log(temp)

  }




    useEffect(() => {
        const pathArray = window.location.pathname.split('/');
        const lastPath = pathArray[pathArray.length - 1];
        setLastDirectory(lastPath);
        data.forEach(element => {
            if(decodeURIComponent(lastDirectory.replace(/\+/g, ' '))==element.name){
                setHotel(element)
            }})
    });


    const getHotelData = (id) => {
      fetch(config.HOST + ":" + config.PORT + "/hotel/" + id)
      .then(response => response.json())
      .then(hotel => {
        setHotel(hotel);
            });
  }

    return (
        <div style={{marginTop:"30px",width:"100vw",height:"100%",display:"flex",alignContent:"center",justifyContent:"center"}}>
            <div style={{marginTop:"30px",width:"900px",height:"1200px", padding:"40px", borderRadius:"10px", backgroundColor:"white"}}>
                <div style={{width:"100%",fontSize:"40px"}}>{hotel.name}</div>
                <div style={{padding:"10px"}}>{hotel.description}</div>
                <img src={hotel.image} style={{width:"100%"}}/> 
                <div style={{border:" 1px solid gray", borderRadius:"10px",paddingTop:"20px",marginTop:"20px"}}>
                  <Calendar
                    style={{width:"100%"}}
                    selected={selectedDates}
                    onChange={handleChange}
                    onOverbook={(e, err) => alert(err)}
                    components={{
                        DayCellFooter: ({ innerProps }) => (
                             <div {...innerProps}></div>
                                 ),
                             }}
                             disabled={(date, state) => !state.isSameMonth}
                             reserved={reserved}
                             reserverFooter={reserved}
                             dateFnsOptions={{ weekStartsOn: 1 }}
                             range={true}
                             isStart={true}
                  />
                     <button style={{marginTop:"30px",width:"10%",position:"relative",bottom:"10px",left:"45%" }} onClick={reserveDate}> Reservar </button>
                 </div>
            </div>
        </div>
    );
  }
  export default Hotel
