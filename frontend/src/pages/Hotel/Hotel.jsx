import { useEffect, useState } from 'react'; 
import {config} from '../../config'
import axios from "axios";
import Calendar from "@demark-pro/react-booking-calendar";
import Modal from '../../components/Modal/Modal';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';




function Hotel() {
    const [lastDirectory, setLastDirectory] = useState('');
    const [hotel, setHotel] = useState({});
    const [modalVisible, setModalVisible] = useState(false);
    const [selectedDates, setSelectedDates] = useState([]);//mandar al apretar reservar
    const handleChange = (e) => {setSelectedDates(e);}
    const [data, setData] = useState([]);
    const [userData, setUserData] = useState(() => {
      const saved = localStorage.getItem("userData");
      const initialValue = JSON.parse(saved);
      console.log(initialValue);
      return initialValue || "";
    });
    const[reserved, setReserved] = useState([
        {
          startDate: new Date(2023, 7, 10),
          endDate: new Date(2023, 7, 15),
        },
      ])

    useEffect(() => {
          const pathArray = window.location.pathname.split('/');
          const lastPath = pathArray[pathArray.length - 1];
          setLastDirectory(lastPath);
          fetch('http://localhost:8090/hotels')
          .then(response => {
            // Check if the request was successful (status code 200)
            if (!response.ok) {
              throw new Error(`HTTP error! Status: ${response.status}`);
            }
        
            // Parse the response body as JSON
            return response.json();
          })
          .then(data => {
            data.forEach(element => {

              console.log(decodeURIComponent(lastPath))
              if(decodeURIComponent(lastPath)==element.name){
                  setHotel(element)
                  console.log(element)

              }})
          })
          .catch(error => {
            // Handle any errors that occurred during the fetch
            console.error('Error:', error);
          });
        

            }, [])


      useEffect( () => {
        getLoginData();
    }, [])

    function funcToggle(){
      setModalVisible(false)
    }

    const getLoginData = () => {
        fetch(config.HOST + ":" + config.PORT + "/login")
        .then(response => response.json())
        .then(data => {
            setData(data);
        });
    }

// desde el back, guardar y devolver las fechas reservadas asi:
//un objeto asi se manda para reservarlo, las reservas para bloquearlas se tienen que mandar en un array de estos. Al guardar una reserva nueva, añadirla al fin del array

async function reserveDate(){
  if (userData) {
    var formValues = {
      date_from: selectedDates[0],
      date_to: selectedDates[1],
      precio: 6,
      user_id: userData.user_id,
      hotel_id: hotel.id,
    };

    try {
      // Show "Your reservation is in progress" toast
      const progressToastId = toast.promise(
        axios.post("http://localhost:8090/reserva", formValues)
          .then((res) => {
            console.log(res);
            // You can return any value you want to use after the toast is resolved
            return 'Reservation successful';
          })
          .catch((err) => {
            console.log(err);
            throw new Error('Reservation failed. Please try again.');
          }),
        {
          pending: 'Tu reserva esta en progreso.',
          success: 'Se ha realizado la reserva exitosamente.',
          error: 'La reserva ha fallado.',
        }
      );

      // Wait for the toast to be resolved
      const result = await toast.promise(progressToastId);
      console.log('Toast resolved with result:', result);
    } catch (error) {
      console.error('Error:', error);
    }
  } else {
    setModalVisible(true);
  }
};





  
 const getHotelData = (id) => {
      fetch(config.HOST + ":" + config.PORT + "/hotel/" + id)
      .then(response => response.json())
      .then(hotel => {
        setHotel(hotel);
            });
  }

    return (
      <div>
        <ToastContainer />
        {modalVisible&&<Modal funcToggle={funcToggle}/>}
        <div style={{width:"100%",height:"100%",display:"flex",alignContent:"center",justifyContent:"center"}}>
            <div style={{marginTop:"30px",width:"900px",height:"1200px", padding:"40px", borderRadius:"10px", backgroundColor:"white"}}>
                <div style={{width:"100%",fontSize:"40px"}}>{hotel.name}</div>
                <div style={{padding:"10px"}}>{hotel.description}</div>
                <div style={{padding:"10px"}}>Amenities: {hotel.Amenities}</div>
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
      </div>
    );
  }
  export default Hotel
