// Obtener el formulario
const formulario = document.getElementById("formulario-reserva");

// evento de escucha para el envío del formulario
formulario.addEventListener("submit", function (event) {
    event.preventDefault();

    // Obtener los datos del formulario
    const nombre = document.getElementById("nombre").value;
    const email = document.getElementById("email").value;
    const hotel = document.getElementById("hotel").value;
    const fechaEntrada = document.getElementById("fecha_entrada").value;
    const fechaSalida = document.getElementById("fecha_salida").value;

    // Validar y procesar la reserva
    const habitacionesDisponibles = verificarDisponibilidadHabitaciones(hotel, fechaEntrada, fechaSalida);

    if (habitacionesDisponibles > 0) {
        const costoTotal = calcularCostoTotal(hotel, fechaEntrada, fechaSalida);
        const reservaId = almacenarReserva(nombre, email, hotel, fechaEntrada, fechaSalida, costoTotal);
        enviarCorreoConfirmacion(email, reservaId);
        alert("¡La reserva se ha realizado con éxito!");
    } else {
        alert("Lo sentimos, no hay habitaciones disponibles para las fechas seleccionadas.");
    }

    // página de confirmación o agradecimiento
     window.location.href = "confirmacion.html";
});

function verificarDisponibilidadHabitaciones(hotel, fechaEntrada, fechaSalida) {
    //lógica para verificar la disponibilidad de habitaciones en el hotel y fechas seleccionados


    return Math.floor(Math.random() * 6); // como ejemplo retorno aqui un numero aleatorio entre 0 y 5
}

function calcularCostoTotal(hotel, fechaEntrada, fechaSalida) {
    // lógica para calcular el costo total de la reserva
    // De forma simplificada, aquí se retorna un costo fijo de $100 por noche de estancia

    const diasEstancia = (new Date(fechaSalida) - new Date(fechaEntrada)) / (1000 * 60 * 60 * 24);
    const costoPorNoche = 100;

    return diasEstancia * costoPorNoche;
}

function almacenarReserva(nombre, email, hotel, fechaEntrada, fechaSalida, costoTotal) {
    // lógica para almacenar la reserva en una base de datos

}

function generateReservaId() {
    // Generar un ID de reserva aleatorio
    // De forma simplificada, aquí se genera un ID basado en la fecha y hora actual

    const timestamp = new Date().getTime();
    const random = Math.floor(Math.random() * 1000);

    return `RES${timestamp}${random}`;
}

function enviarCorreoConfirmacion(email, reservaId) {
    // Aquí iría la lógica para enviar un correo electrónico de confirmación


    const mensaje = `Gracias por realizar la reserva. Tu ID de reserva es: `;

    // Código para enviar el correo electrónico...
}