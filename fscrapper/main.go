package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func turnToDust(root string, substr string) error {

	counter := 0
	err := filepath.Walk(root,
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(path, substr) {
				fmt.Printf("Menghapus file: %s\n", path)
				err = os.Remove(path)
				if err != nil {
					return err
				}
				counter += 1
			}
			return nil
		})
	if err != nil {
		return err
	}
	fmt.Printf("Berhasil menghapus %d file\n", counter)
	return nil

}

func main() {
	fmt.Println("Masukkan root directory untuk memulai scanning:")
	var dst_path string
	var ext_target string
	fmt.Scanln(&dst_path)
	fmt.Printf("Direktori yang Anda masukkan adalah: %s\n", dst_path)
	fmt.Println("Masukkan extension file yang ingin dihapus:")
	fmt.Scanln(&ext_target)
	fmt.Printf("Ekstensi file yang Anda masukkan adalah: %s\n", ext_target)
	fmt.Println("Memulai proses scanning...")
	err := turnToDust(dst_path, ext_target)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Berhasil menghapus semua file yang diinginkan.")
}
