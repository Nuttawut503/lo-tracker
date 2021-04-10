FROM cirrusci/flutter:2.0.4
WORKDIR /root
RUN apt-get update \
	&& curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
	&& apt-get install -y ./google-chrome-stable_current_amd64.deb \
	&& rm google-chrome-stable_current_amd64.deb
