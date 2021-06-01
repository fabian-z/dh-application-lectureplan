// ES6 module imports here
import 'bootstrap';

// Webpack CSS imports
// Could use Bootstrap SCSS sources here
// see https://getbootstrap.com/docs/5.0/getting-started/webpack/#importing-styles
// precompiled CSS works for now
import 'bootstrap/dist/css/bootstrap.css';

import "../css/main.css";

// npm install @fullcalendar/core @fullcalendar/daygrid @fullcalendar/timegrid @fullcalendar/list
import { Calendar } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import listPlugin from '@fullcalendar/list';
import interactionPlugin from '@fullcalendar/interaction';
import bootstrapPlugin from '@fullcalendar/bootstrap'
import scrollgridPlugin from '@fullcalendar/scrollgrid'

let calendar = new Calendar(calendarEl, {
  plugins: [ dayGridPlugin, timeGridPlugin, listPlugin, interactionPlugin, bootstrapPlugin, scrollgridPlugin ],
  initialView: 'dayGridMonth',
  headerToolbar: {
    left: 'prev,next today',
    center: 'title',
    right: 'dayGridMonth,timeGridWeek,listWeek'
  }
});


document.addEventListener('DOMContentLoaded', function() {
    // Setup handlers when DOM is ready

    // DOM event handlers

    document.getElementById("submit").addEventListener("click", function() {
       // Example click handler
    }, false);

}, false);
