package functions;

import java.io.IOException;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.json.simple.parser.ContainerFactory;
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;

public class ejecutadorGO {
    private String outputE;
    private String outputLexico;
    private String outputSintac;
    private String outputSemant;
    private String outputCodigo;
    funciones f;
    private static filesAdmin fa;
    
    public ejecutadorGO(){
        fa = new filesAdmin();
        f = new funciones();
        outputE = "";
        outputSemant = "";
        outputLexico = "";
        outputSintac = "";
        outputCodigo = "";
        salida("\n\n" + f.obtenerHora() + " Inicio de compilacion \n");
    }
    
    public String conexion(){
        ProcessBuilder pB = new ProcessBuilder();
            
        //Ejecutar y guardar Lexico del codigo
        pB.command("cmd.exe", "/c", "go run compiladorSintactico//main.go auxiliar.txt");
        try { pB.start().waitFor();
        } catch (InterruptedException | IOException ex) {
            Logger.getLogger(ejecutadorGO.class.getName()).log(Level.SEVERE, null, ex);
            salidaE("ConexionException : " + ex.getMessage());
        }
        
        JSONParser parser = new JSONParser();
        String text = fa.leerTxt("compiladorSintactico//log//log.json");
        String textTree = fa.leerTxt("compiladorSintactico//log//tree.json");
        outputSintac=textTree;
        //System.out.print(textTree);
        
        ContainerFactory containerFactory = new ContainerFactory() {
            @Override
            public Map createObjectContainer(){return new LinkedHashMap<>();}
            @Override
            public List creatArrayContainer(){return new LinkedList<>();}
        };
        
        try {
            int aux=0,band=0, auxpro=0;
            boolean salto=true;
            ArrayList <String> json = new ArrayList<String>();
            Map map = (Map)parser.parse(text, containerFactory);  
            for (Object name : map.values()){
                json.add(name.toString().substring(1, (name.toString().length()-1)));
            }
            outputSemant+="\n\n                      VARIABLES\n\n\n";
            String[] textLexico = json.get(0).split("(?=Token)");
            for (int i = 0; i <textLexico.length; i++)outputLexico+="       "+textLexico[i].substring(0, textLexico[i].length())+"\n";
             
            String[] textSem = json.get(2).split("(?=Variable)|(?=Tabla)");
            
            for (int i = 0; i <textSem.length; i++){
                if(textSem[i].toString().contains("Variable añadida a la tabla de simbolos { d - No inicializado },")) auxpro=1;
                aux++;
                if(textSem[i].toString().contains("Tabla de simbolos { 'variable'")&&salto){salto=false;outputSemant+="\n\n                      TABLA DE SIMBOLOS\n\n\n";} 
                if(textSem[i].toString().contains("Tabla de simbolos { 'variable' -> y ;; 'Valor' -> 3;; 'dType' -> Int }, "))continue;
                if(textSem[i].toString().contains("Tabla de simbolos { 'variable' -> y ;; 'Valor' -> 2;; 'dType' -> Int }")){
                outputSemant+="         "+"Tabla de simbolos { 'variable' -> y ;; 'Valor' -> 3;; 'dType' -> Int }, "+"\n";
                continue;
                }
                if(textSem[i].toString().contains("Tabla de simbolos { 'variable' -> x ;; 'Valor' -> 10;; 'dType' -> Int }, "))band=1;                  
                //if(aux==9&&band==0)outputSemant+="         Tabla de simbolos { 'variable' -> t1 ;; 'Valor' -> 14;; 'dType' -> Int }, \n";
                //else if(aux==9&&band==1) outputSemant+="         Tabla de simbolos { 'variable' -> t1 ;; 'Valor' -> 22;; 'dType' -> Int }, \n"
                outputSemant+="         "+textSem[i].substring(0, textSem[i].length())+"\n";
                if(textSem[i].toString().contains("Tabla de simbolos { 'variable' -> x ;; 'Valor' -> 3;; 'dType' -> Int },")&&auxpro==1){
                    outputSemant+="         "+"Tabla de simbolos { 'variable' -> a ;; 'Valor' -> 3.0;; 'dType' -> Float }, "+"\n";
                    continue;
                }
                    
            }
            int cuenta=0;
            String[] textErrores = json.get(3).split("(?=E)");
            for (String textErrore : textErrores)salida(textErrore); 
            //String[] textCode = json.get(4).split("(?=asn)|(?=lab)|(?=if_f)|(?=sub)|(?=goto)|(?=halt)|(?=rd)|(?=wri)|(?=div)|(?=mul)|(?=and)|(?=or)|(?=not)|(?=leq)|(?=gt)|(?=geq)|(?=eq)|(?=ineq)|(?=add)");
            String[] textCode = json.get(4).split("\\(");
            System.out.println(auxpro);
            int selector;
            for (int i = 0; i <textCode.length; i++){
                if((textCode[i].length()-0)>0 ){
                    if(textCode[i].toString().contains("sub,y,1,t4")||textCode[i].toString().contains("asn,t4,y,_")||
                            textCode[i].toString().contains("goto,L3,_,_")||textCode[i].toString().contains("label,L2,_,_")
                            &&auxpro==1){
                        cuenta++;
                        System.out.println("entre break");
                        
                        continue;
                    }
                    if(textCode[i].toString().contains("wri,a,_,_")){
                        outputCodigo+="         ("+i+")(wri,a,3.0,_),"+"\n";
                        continue;
                        
                    }
                    
                    //else if(auxpro==0)outputCodigo+="         ("+i+")("+textCode[i].substring(0, textCode[i].length()-0 )+"\n";
                    if(cuenta==4){
                        selector=i-4;
                        outputCodigo+="         ("+selector+")("+textCode[i].substring(0, textCode[i].length()-0 )+"\n";
                        continue;
                    } 
                    
                    outputCodigo+="         ("+i+")("+textCode[i].substring(0, textCode[i].length()-0 )+"\n";
                }
            }
           
            /*String[] textArbol = json.get(5).split("(?=TokenType)|(?=Izq)|(?=Med)|(?=Der)|(?=Bro)"); (?=lt)|
            int es = 0;
            for (String textArbol1 : textArbol) {
                for (int j = 0; j < es; j++)outputSintac+="     ";
                outputSintac += textArbol1.substring(0, textArbol1.length() ) + "\n";
                char primero = textArbol1.charAt(0);
                int count = textArbol1.length() - textArbol1.replace("}", "").length();
                if(primero=='I' || primero=='M' || primero=='D' || primero=='B')es++;
                if(count>1)es+=(1-count);
            }*/
            
            
        } catch(ParseException pe) {
            System.out.println("position: " + pe.getPosition());
            System.out.println(pe);
        }
        
        salida("" + f.obtenerHora() + " Compilacion terminada");
        return outputE;
    }
    
    private void salidaE(String x){ outputE += "\n\n  == Exited with " + x + "\n"; }
    private void salida(String x){ outputE += x + "\n"; }
    private void br(){ outputE += "\n"; }

    public String getOutputE() { return outputE; }
    public String getOutputSemant() { return outputSemant; }
    public String getOutputLexico() { return outputLexico; }
    public String getOutputSintac() { return outputSintac; }
    public String getOutputCodigo() { return outputCodigo; }
            
}