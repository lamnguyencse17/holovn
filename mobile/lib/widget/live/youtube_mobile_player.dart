import 'package:flutter/material.dart';
import 'package:youtube_player_flutter/youtube_player_flutter.dart';


class YoutubeMobilePlayer extends StatefulWidget {
  final String liveId;
  YoutubeMobilePlayer({Key? key, required this.liveId}) : super(key: key);

  @override
  _YoutubeMobilePlayerState createState() => _YoutubeMobilePlayerState(this.liveId);
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
  _YoutubeMobilePlayerState(this.liveId);

  @override
  void initState() {
    super.initState();
    _controller = YoutubePlayerController(
      initialVideoId: liveId,
      flags: YoutubePlayerFlags(
        mute: false,
        autoPlay: true,
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
      return YoutubePlayer(
        controller: _controller,
        width: constraints.maxWidth,
        showVideoProgressIndicator: true,
        onReady: () {
          print('Player is ready.');
        },
      );
    });
  }
}
