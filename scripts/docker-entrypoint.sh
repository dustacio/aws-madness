#!/bin/sh

[ -z $CLOUDFRONT_URL ] && echo "CLOUDFRONT_URL is Not Set" && exit

/aws-madness
