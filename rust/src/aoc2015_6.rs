use utils;

#[derive(Clone, Debug)]
struct Light {
    x: i32,
    y: i32,
    lit: bool,
}

impl Light {
    fn is_lit(&self) -> bool {
        self.lit
    }
}

#[derive(Clone)]
struct LightGrid {
    width: i32,
    height: i32,
    lights: Vec<Light>,
}

impl LightGrid {
    // maximum possible size of the grid
    fn max_size(&self) -> i32 {
        self.width * self.height
    }

    fn add(&mut self, i: Light) {
        &self.lights.push(i);
    }

    // returns the current size of the grid, ie number of elements in there
    fn size(&self) -> usize {
        self.lights.len()
    }

    // fill the grid with default lights
    fn fill(&mut self) {
        let mut i: i32 = 0;
        let size = self.max_size();
        let w = self.width;
        while i < size {
            let x = i % w;
            // integer division for y pos
            let y = i / w;
            self.add(Light { x, y, lit: false });
            i += 1;
        }
    }

    fn items(self) -> Vec<Light> {
        self.lights
    }
}

#[test]
fn will_fill_the_grid_with_lights_in_default_state() {
    let width = 5;
    let height = 5;
    let mut grid = LightGrid {
        width,
        height,
        lights: Vec::new(),
    };

    &grid.fill();
    // iterate over each item in the grid and check the position + the light status
    for i in grid.items().iter() {
        // println!("{:?}", i);
        assert_eq!(false, i.is_lit());
    }
}

fn calculate_lights(data: &str) -> i32 {
    0
}

pub fn solve() -> String {
    let data = utils::read_file("../data/2015_6.txt");
    let lights_lit = calculate_lights(&data);
    format!("{:?} christmas lights have been lit", &lights_lit)
}
