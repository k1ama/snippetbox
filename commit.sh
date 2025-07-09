chmod -R 777 ../snippetbox
rm -r -f test/log

echo "------ SYNC LOCAL COPY WITH RECENT REMOTE CHANGES..."
git pull
echo "------ DONE"

echo "------ PREPARING LOCAL PUSH PACKAGE..."
git add .
git commit
echo "------ DONE"

echo "------ UPLOADING CHANGES"
git push
echo "------ COMPLETED"
