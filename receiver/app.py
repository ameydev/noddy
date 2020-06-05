# http://flask.pocoo.org/docs/patterns/fileuploads/
import os
from flask import Flask, request, redirect, url_for, send_from_directory
from werkzeug import secure_filename
import tarfile


UPLOAD_FOLDER = '/workspace'
# ALLOWED_EXTENSIONS = set(['tar', 'pdf', 'png', 'jpg', 'jpeg', 'gif'])
ALLOWED_EXTENSIONS = set(['tar','.gz','tar.gz'])

app = Flask(__name__)
app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

def un_tar(file_path):
    my_tar = tarfile.open(file_path)
    my_tar.extractall(str(file_path).split('.tar')[0])
    my_tar.close()
    print('removing the tar')
    os.remove(file_path)

def allowed_file(filename):
  # this has changed from the original example because the original did not work for me
    return filename[-3:].lower() in ALLOWED_EXTENSIONS

@app.route('/', methods=['GET', 'POST'])
def upload_file():
    if request.method == 'POST':
        file = request.files['file']
        print(file.filename)
        print(str(allowed_file(file.filename)))
        if file and allowed_file(file.filename):
            print('**found file', file.filename)
            filename = secure_filename(file.filename)
            file.save(os.path.join(app.config['UPLOAD_FOLDER'], filename))
            print('Extracting the image tar..')
            un_tar(os.path.join(app.config['UPLOAD_FOLDER'], filename))
            return url_for('uploaded_file',
                                    filename=filename)
    return '''
    <!doctype html>
    <title>Upload new File</title>
    <h1>Upload new File</h1>
    <form action="" method=post enctype=multipart/form-data>
      <p><input type=file name=file>
         <input type=submit value=Upload>
    </form>
    '''

@app.route('/uploads/<filename>')
def uploaded_file(filename):
    return send_from_directory(app.config['UPLOAD_FOLDER'],
                               filename)

@app.route('/')
def hello():
    return "hello world" 
           


if __name__ == '__main__':
	app.run(host='0.0.0.0')