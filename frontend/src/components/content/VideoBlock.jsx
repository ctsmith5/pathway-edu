import React from 'react';

const VideoBlock = ({ data }) => {
  if (!data?.url) {
    return null;
  }

  const { url, title = 'Video' } = data;

  // Convert YouTube URLs to embed format
  const getEmbedUrl = (videoUrl) => {
    // YouTube watch URLs
    const youtubeMatch = videoUrl.match(/(?:youtube\.com\/watch\?v=|youtu\.be\/)([^&\s]+)/);
    if (youtubeMatch) {
      return `https://www.youtube.com/embed/${youtubeMatch[1]}`;
    }
    
    // Vimeo URLs
    const vimeoMatch = videoUrl.match(/vimeo\.com\/(\d+)/);
    if (vimeoMatch) {
      return `https://player.vimeo.com/video/${vimeoMatch[1]}`;
    }

    // Return as-is if already an embed URL or unknown format
    return videoUrl;
  };

  const embedUrl = getEmbedUrl(url);

  return (
    <div className="video-block">
      <div className="video-block__container">
        <iframe
          src={embedUrl}
          title={title}
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          allowFullScreen
          className="video-block__iframe"
        />
      </div>
      {title && <p className="video-block__title">{title}</p>}
    </div>
  );
};

export default VideoBlock;

