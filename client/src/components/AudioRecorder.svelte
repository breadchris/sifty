<script lang="ts">
  import {Button} from "flowbite-svelte";
  import {clips} from "../store/content";
  import {blobToUint8Array} from "../util/blob";

  let stream;
  let mediaRecorder;
  let recordButton;
  let chunks = [];

  async function onStop(e) {
    console.log("recorder stopped");
    const blob = new Blob(chunks, { type: "audio/mp3" });
    chunks = [];
    const audioURL = window.URL.createObjectURL(blob);
    clips.set([
      ...$clips,
      {
        src: audioURL,
        content: await blobToUint8Array(blob)
      },
    ]);
  }

  async function setup() {
    if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
      console.log("getUserMedia supported.");
      stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      mediaRecorder = new MediaRecorder(stream, {
        mimeType: "audio/webm;codecs=opus",
        audioBitsPerSecond: 96000,
      });
	  mediaRecorder.ondataavailable = (e) => {
		  chunks.push(e.data);
	  }
      mediaRecorder.onstop = onStop;
    }
  }

  setup();
  function record() {
    console.log("recorder started");
    mediaRecorder.start();
  }
  function stop() {
    mediaRecorder.stop();
    console.log(mediaRecorder.state);
    console.log("recorder stopped");
  }
</script>

<div class="wrapper">
  <section class="main-controls">
    <canvas class="visualizer" height="60px" />
    <div id="buttons">
      <Button class="record" on:click={record} bind:this={recordButton}>
          Record
      </Button>
      <Button class="stop" on:click={stop}>Stop</Button>
    </div>
  </section>

  <section class="sound-clips">
    {#each $clips as clip}
      <article class="clip">
        <audio src={clip.src} controls/>
        <Button>Delete</Button>
      </article>
    {/each}
  </section>
</div>

<style>
  .main-controls {
    padding: 0.5rem 0;
  }
  canvas {
    display: block;
    margin-bottom: 0.5rem;
  }
  /* Make the clips use as much space as possible, and
 * also show a scrollbar when there are too many clips to show
 * in the available space */
  .sound-clips {
    flex: 1;
    overflow: auto;
  }
  .clip {
    padding-bottom: 1rem;
  }
  audio {
    width: 100%;
    display: block;
    margin: 1rem auto 0.5rem;
  }
</style>
