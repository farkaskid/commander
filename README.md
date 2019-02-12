# commander
A small go program that can be used to voice control mundane tasks on your PC.

This implementation is very specific to linux as I've used [`pulseaudio`](https://www.freedesktop.org/wiki/Software/PulseAudio/) API for picking up audio from microphone.

On linux, Your sound server may be different. `pulseaudio` is the default sound server on many beginner distros but you can also go for something like ALSA or JACK.

For Windows and MacOS, [portaudio](http://www.portaudio.com/) is recommended.

I've used [PocketSphinx](http://www.speech.cs.cmu.edu/pocketsphinx/) for getting the text according to the audio, however `Sphinx4` also works well, it's just a little slow. The go bindings can be found [here](https://github.com/cmusphinx/pocketsphinx)

I've used and highly recommend a custom dictionary according to your desirable commands which then can be mapped to anything you want. This dictionary is very small and only consists of few commands, more commands should be added to improve the accuracy.

This is just a basic working prototype and a lot more will be added.

### Supported actions
- Lock the system.
- Unlock the system.(without password with voice)

New actions can be very easily added with a big enough dictionary.
