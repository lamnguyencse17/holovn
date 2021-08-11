import 'package:flutter/material.dart';
import 'package:youtube_player_flutter/youtube_player_flutter.dart';


class YoutubeMobilePlayer extends StatefulWidget {
  final String liveId;
  final String status;
  YoutubeMobilePlayer({Key? key, required this.liveId, required this.status}) : super(key: key);

  @override
  _YoutubeMobilePlayerState createState() => _YoutubeMobilePlayerState(this.liveId, this.status);
}

class _YoutubeMobilePlayerState extends State<YoutubeMobilePlayer> {
  double _volume = 100;
  bool _muted = false;
  bool _isPlayerReady = false;

  late YoutubePlayerController _controller;
  late TextEditingController _idController;
  late TextEditingController _seekToController;

  late PlayerState _playerState;
  late YoutubeMetaData _videoMetaData;

  final String liveId;
  final String status;
  _YoutubeMobilePlayerState(this.liveId, this.status);

  @override
  void initState() {
    super.initState();
    _controller = YoutubePlayerController(
      initialVideoId: liveId,
      flags: YoutubePlayerFlags(
        isLive: this.status == "live" ? true : false,
        hideControls: false,
        mute: false,
        autoPlay: true,
        useHybridComposition: true
      ),
    );
    _controller.addListener(listener);
    _idController = TextEditingController();
    _seekToController = TextEditingController();
    _videoMetaData = const YoutubeMetaData();
    _playerState = PlayerState.unknown;
  }

  @override
  void deactivate() {
    // Pauses video while navigating to next page.
    _controller.pause();
    super.deactivate();
  }

  @override
  void dispose() {
    _controller.dispose();
    _idController.dispose();
    _seekToController.dispose();
    super.dispose();
  }

  void listener() {
    if (_isPlayerReady && mounted && !_controller.value.isFullScreen) {
      setState(() {
        _playerState = _controller.value.playerState;
        _videoMetaData = _controller.metadata;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
        builder: (BuildContext context, BoxConstraints constraints) {
      return YoutubePlayerBuilder(player: YoutubePlayer(
        controller: _controller,
        width: constraints.maxWidth,
        showVideoProgressIndicator: true,
        onReady: () {
          print('Player is ready.');
        },
      ), builder: (context, player){
        return Column(children: [player],);
      });

    });
  }
}
