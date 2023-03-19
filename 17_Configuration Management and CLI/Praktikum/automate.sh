#!/bin/bash

# Fungsi untuk membuat direktori dan file-file yang dibutuhkan
function create_files() {
  folder_name="$1"

  # Membuat folder baru dengan format tanggal dan waktu saat ini
  current_time=$(date "+%Y-%m-%d %H:%M:%S")
  mkdir "$folder_name at $current_time"

  # Membuat file about_me
  mkdir "$folder_name at $current_time/about_me"
  mkdir "$folder_name at $current_time/about_me/personal"
  echo "https://www.facebook.com/$2/" > "$folder_name at $current_time/about_me/personal/facebook.txt"
  mkdir "$folder_name at $current_time/about_me/professional"
  echo "https://www.linkedin.com/in/$3/" > "$folder_name at $current_time/about_me/professional/linkedin.txt"

  # Membuat file my_friends
  mkdir "$folder_name at $current_time/my_friends"
  curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%20of%20my%20friends.txt > "$folder_name at $current_time/my_friends/list_of_my_friends.txt"

  # Membuat file my_system_info
  mkdir "$folder_name at $current_time/my_system_info"
  echo "My username: $(whoami)" > "$folder_name at $current_time/my_system_info/about_this_laptop.txt"
  echo "$(uname -a)" >> "$folder_name at $current_time/my_system_info/about_this_laptop.txt"
  ping google.com -n 3 > "$folder_name at $current_time/my_system_info/internet_connection.txt"
}

# Fungsi utama
function main() {
  # Cek apakah argumen cukup
  if [ $# -ne 3 ]; then
    echo "Usage: ./automate.sh <nama_folder> <facebook_username> <linkedin_username>"
    exit 1
  fi

  # Memanggil fungsi create_files
  create_files "$1" "$2" "$3"

  # Menampilkan pesan sukses
  echo "Files and directories have been created successfully."
  echo "Folder name: $1 at $(date "+%Y-%m-%d %H:%M:%S")"
}

# Panggil fungsi utama dengan argumen dari baris perintah
main "$@"

