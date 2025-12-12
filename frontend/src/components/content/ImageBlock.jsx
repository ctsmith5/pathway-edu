import React, { useState } from 'react';

const ImageBlock = ({ data }) => {
  const [error, setError] = useState(false);

  if (!data?.url) {
    return null;
  }

  const { url, alt = '', caption = '' } = data;

  if (error) {
    return (
      <div className="image-block image-block--error">
        <div className="image-block__placeholder">
          <span>Image not available</span>
        </div>
        {caption && <p className="image-block__caption">{caption}</p>}
      </div>
    );
  }

  return (
    <figure className="image-block">
      <img 
        src={url} 
        alt={alt}
        className="image-block__img"
        onError={() => setError(true)}
        loading="lazy"
      />
      {caption && (
        <figcaption className="image-block__caption">{caption}</figcaption>
      )}
    </figure>
  );
};

export default ImageBlock;

