echo '== prepare build =='

rm -rf app
mkdir app
mkdir -p app/static/upload/images

cp -R conf app
cp main app

echo '== start build =='

docker build -t phoenix/sisyphus:v1.0 .
