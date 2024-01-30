
use image::{ImageBuffer, Rgb};
use num_complex::Complex;

fn main() {
    let width = 1920;
    let height = 1080;

    let mut img = ImageBuffer::new(width, height);

    let scalex = 3.0 / width as f32;
    let scaley = 3.0 / height as f32;

    for (x, y, pixel) in img.enumerate_pixels_mut() {
        let cx = y as f32 * scalex - 1.5;
        let cy = x as f32 * scaley - 1.5;

        let c = Complex::new(cx, cy);
        let mut z = Complex::new(0.0, 0.0);

        let mut i = 0;
        for t in 0..255 {
            if z.norm() > 2.0 {
                break;
            }
            z = z * z + c;
            i = t;
        }

        *pixel = Rgb([i as u8, (255 - i) as u8, (i / 2) as u8]);
    }

    img.save("fractal.png").unwrap();
    println!("Fractal generated and saved as 'fractal.png'");
}
