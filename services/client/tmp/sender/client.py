import requests,sys
import tarfile
import os.path

url = "http://localhost:5000/"

dir_path= sys.argv[1]
tar_name = 'sr.tar.gz'


def make_tarfile(output_filename, source_dir):
    with tarfile.open(output_filename, "w:gz") as tar:
        tar.add(source_dir, arcname=os.path.basename(source_dir))


make_tarfile(tar_name, dir_path)
fin = open(tar_name, 'rb')
files = {'file': fin}
try:
  r = requests.post(url, files=files)
  print(r.text)
finally:
	fin.close()


