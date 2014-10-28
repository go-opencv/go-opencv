FROM ubuntu:trusty

RUN echo "deb http://ppa.launchpad.net/jon-severinsson/ffmpeg/ubuntu trusty main" >> /etc/apt/sources.list

RUN echo 'deb http://archive.ubuntu.com/ubuntu trusty main' >> /etc/apt/sources.list

RUN echo 'deb http://archive.ubuntu.com/ubuntu trusty multiverse' >> /etc/apt/sources.list

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1DB8ADC1CFCA9579

RUN apt-get update

RUN apt-get install -q -y frei0r-plugins

RUN apt-get install -y -q python-pip

RUN apt-get install -y -q wget curl unzip

RUN apt-get install -y -q build-essential

RUN apt-get install -y -q cmake

RUN apt-get install -y -q python2.7 python2.7-dev

#RUN wget 'https://pypi.python.org/packages/2.7/s/setuptools/setuptools-0.6c11-py2.7.egg' && /bin/sh setuptools-0.6c11-py2.7.egg && rm -f setuptools-0.6c11-py2.7.egg

RUN pip install numpy

RUN apt-get install -y -q libavformat-dev libavcodec-dev libavfilter-dev libswscale-dev

RUN apt-get install -y -q libjpeg-dev libpng-dev libtiff-dev libjasper-dev zlib1g-dev libopenexr-dev libxine-dev libeigen3-dev libtbb-dev

RUN apt-get install -y libopencv-dev checkinstall pkg-config yasm libtiff4-dev libjpeg-dev libjasper-dev libavcodec-dev libavformat-dev libswscale-dev libdc1394-22-dev libxine-dev libgstreamer0.10-dev libgstreamer-plugins-base0.10-dev libv4l-dev python-dev python-numpy libtbb-dev libqt4-dev libgtk2.0-dev libfaac-dev libmp3lame-dev libopencore-amrnb-dev libopencore-amrwb-dev libtheora-dev libvorbis-dev libxvidcore-dev x264 v4l-utils ffmpeg

RUN apt-get update -y && apt-get install --no-install-recommends -y -q curl build-essential ca-certificates git mercurial bzr

RUN mkdir /goroot && curl https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

RUN mkdir /gopath

ENV GOROOT /goroot

ENV GOPATH /gopath

ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
