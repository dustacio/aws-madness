# AWS S3 Filename Corrector

AWS S3 encodes the path portion of files stored in buckets.
S3 can't be used as a mirror for things like RPMs because
filenames like

   /x86_64/libstdc++-devel-32bit-4.3-62.200.2.x86_64.rpm

need to be requested as

   /x86_64/libstdc%2B%2B-devel-32bit-4.3-62.200.2.x86_64.rpm

This is a simple proxy that will redirect requests to the proper encoded URL.

This issue is discussed here:
    https://forums.aws.amazon.com/thread.jspa?messageID=722673

## Getting Started

`go build`

### Required Environment Variables

* `HOST` default is "", the host to listen on
* `PORT` default is "8000", the port to listen on
* `CLOUDFRONT_URL`: no default, required, the full URL to your Cloud Front URL, e.g. `https://example.com` (no trailing slash)

## Deployment

Make a docker image, wire up to your URL, redirect to your Cloud Front URL

## Author

Dusty Jones <dusty@dustyjones.com>

## LICENSE

This project is licensed under the MIT License - see the LICENSE.md file for details
