
#include <iostream>
#include <iomanip>
#include <unistd.h>

using namespace std;

void cetakString(string arr[], int size)
{
	for (int i = 0; i < size; i++)
	{
		cout << arr[i] << "	";
	}
	cout << endl;
}

void cetakInt(int arr[], int size)
{
	for (int i = 0; i < size; i++)
	{
		cout << arr[i] << "	";
	}
	cout << endl;
}

int cariBarang(string arr[], int size, string nama)
{
	int index = -1;
	for (int i = 0; i < size; i++)
	{
		if (arr[i] == nama)
		{
			index = i;
		}
	}
	return index;
}

int main()
{

	string nama_barang[5] = {"odol", "sikat", "sabun", "rinso", "mamel"};
	int harga_barang[5] = {5000, 7500, 10000, 6000, 500};
	int index, k = 0, total = 0;
	int uang, kembalian, jumBar[100], harga[100], subTol[100];
	string pilih, barang[100];
	char ulang;
	bool proses = true;

	cetakString(nama_barang, 5);
	cetakInt(harga_barang, 5);

	do
	{
		cout << endl;
		cout << "masukan nama barang : ";
		cin >> pilih;
		index = cariBarang(nama_barang, 5, pilih);
		
		if (index == -1)
		{
			cout << "barang yang anda inputkan tidak tersedia" << endl;
		}
		else
		{
			harga[k] = harga_barang[index];
			barang[k] = nama_barang[index];
			cout << "jumlah barang: ";
			cin >> jumBar[k];
			subTol[k] = jumBar[k] * harga[k];
			total += subTol[k];
			k++;
		}

		cout << "apakah anda ingin membeli yang lain?(y/t): ";
		cin >> ulang;
	} while (ulang == 'y');
	
	// PROSES CETAK STRUK
	while (proses == true)
	{
		system("cls");
		if (k == 0)
		{
			cout << "===TERIMKASIH TELAH MAMPIR KE MINIMARKET KAMI===" << endl;
		}
		else
		{
			cout << "============================================" << endl;
			cout << "\t\tSRTUK BELANJA" << endl;
			cout << "--------------------------------------------" << endl;
			cout << "   BARANG        HARGA          JUMLAH     " << endl;
			cout << "--------------------------------------------" << endl;

			for (int i = 0; i < k; i++)
			{
				cout << "  " << setiosflags(ios::left) << setw(12) << barang[i];
				cout << "	Rp." << setiosflags(ios::left) << setw(12) << harga[i];
				cout << "	 x" << setiosflags(ios::left) << setw(13) << jumBar[i] << endl;
			}
			cout << "--------------------------------------------" << endl;
			cout << " total		Rp." << total << endl;
			cout << " Uang Tunai	Rp.";
			cin >> uang;
			if (uang < total || uang < 0)
			{
				cout << " UANG ANDA KURANG!" << endl;
				cout << "--------------------------------------------" << endl;
				cout << " Ingin melanjutkan pembayaran?(y/t) : ";
				cin >> ulang;
				if (ulang != 'y')
				{
					cout << "--------------------------------------------" << endl;
					cout << "=TERIMKASIH TELAH MAMPIR KE MINIMARKET KAMI=" << endl;
					cout << "============================================" << endl;
					proses = false;
				}
			}
			else
			{
				kembalian = uang - total;
				cout << " Kembalian	Rp." << kembalian << endl;
				cout << "--------------------------------------------" << endl;
				cout << "         TERIMAKSIH TELAH BELANJA" << endl;
				cout << "            DI MINIMARKET KAMI" << endl;
				cout << "============================================" << endl;
				proses = false;
			}
		}
	}

	return 0;
}
