using System;

class URI1047
{
    static void Main()
    {
        var resp = Console.ReadLine().Split(' '); 
        int hi = Convert.ToInt32(resp[0]), mi = Convert.ToInt32(resp[1]), hf = Convert.ToInt32(resp[2]), mf = Convert.ToInt32(resp[3]);
            if ((hf < hi) || (mf<=mi && hf==hi))
                hf += 24;
            hf = (hf*60+mf) - (hi*60+mi);
            Console.WriteLine($"O JOGO DUROU {hf/60} HORA(S) E {hf-(hf/60)*60} MINUTO(S)");
        
    }    
}   