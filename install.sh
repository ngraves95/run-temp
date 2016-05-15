#!/bin/bash

SCRIPTNAME=run-temp-background.sh

go install run-temp.go

TEMPFILE=`tempfile`
crontab -l > $TEMPFILE

echo "@reboot $PWD/$SCRIPTNAME &" >> $TEMPFILE

# Install new crontab
crontab $TEMPFILE
rm $TEMPFILE

cat <<EOF > ./$SCRIPTNAME && chmod +x $SCRIPTNAME
#!/bin/bash

while true; do
    $GOBIN/run-temp 2>/dev/null
    sleep 10
done

EOF
