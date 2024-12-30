Program Liga Bola
Kamus:
    variabel 1,2,3,4 = integer
algoritma
    input user (1 2)
    input user (3 4)

    // Menghitung total menit untuk waktu mulai dan selesai
    mulai = h1 * 60 + m1
    selesai = h2 * 60 + m2

    durasi = selesai - mulai

    if durasi < 0 then
        durasi = durasi + (12 * 60)

    // Menghitung jam dan menit dari durasi
    jam = durasi / 60
    menit = durasi % 60

    output jam dan menit
endprogram

