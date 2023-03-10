
Console.Write("Bir sayı girin: ");
int sayi = int.Parse(Console.ReadLine());

// 4'ün katlarından en yakınını hesapla
int enYakinKat = sayi - (sayi % 4);

// Eğer en yakın 4'ün katı 10'dan küçük veya eşitse, ekrana yazdır
if (enYakinKat <= 10)
{
    Console.WriteLine(enYakinKat);
}
else
{
    Console.WriteLine("Girilen sayıya en yakın 4'ün katı 10'dan büyük.");
}
