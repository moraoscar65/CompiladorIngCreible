/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package functions;

import java.util.Calendar;
import java.util.GregorianCalendar;


public class funciones {

    public funciones() {
    }
    
    // nos da la hora del momento en que se llama a la funcion [ HH : mm : ss ] 
    public String obtenerHora(){
        String tiempo;
        int hora, minutos, segundos;
        Calendar calendario = new GregorianCalendar();
        hora =calendario.get(Calendar.HOUR_OF_DAY);
        minutos = calendario.get(Calendar.MINUTE);
        segundos = calendario.get(Calendar.SECOND);
        
        tiempo = hora+":"+minutos+":"+segundos;
        
        return tiempo;
    }
    
    public int porcentaje(int cantidad, int porcentaje){
        float x = (cantidad*porcentaje)/100;
        int y = Math.round(x);
        return y;
    }
}
