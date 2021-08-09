import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:youtube_player_iframe/youtube_player_iframe.dart';

class YoutubeWebPlayer extends StatefulWidget {
  final String liveId;
  YoutubeWebPlayer({Key? key, required this.liveId}) : super(key: key);

  @override
  _YoutubeWebPlayerState createState() => _YoutubeWebPlayerState(this.liveId);
}

class _YoutubeWebPlayerState extends State<YoutubeWebPlayer> {
  final String liveId;
  late YoutubePlayerController _controller;
  _YoutubeWebPlayerState(this.liveId);

  @override
  void initState(){
    super.initState();
    _controller = YoutubePlayerController(
      initialVideoId: liveId,
      params: const YoutubePlayerParams(
        showControls: true,
        showFullscreenButton: true,
        desktopMode: true,
        privacyEnhanced: true,
        useHybridComposition: true,
      ),
    );
    _controller.onEnterFullscreen = () {
      SystemChrome.setPreferredOrientations([
        DeviceOrientation.landscapeLeft,
        DeviceOrientation.landscapeRight,
      ]);
    };
  }


  @override
  Widget build(BuildContext context) {
    return YoutubePlayerIFrame(controller: _controller,);
  }
}