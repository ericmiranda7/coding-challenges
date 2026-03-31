class Solution {
    public int carFleet(int target, int[] position, int[] speed) {
        List<Car> cars = IntStream
            .range(0, position.length)
            .boxed()
            .map(i -> new Car(position[i], speed[i]))
            .sorted()
            .toList();

        List<Double> finalPos = new LinkedList();

        for (Car car : cars) {
            double fp = (target - car.position) / (double) car.speed;
            car.timeToReach = fp;
        }

        if (cars.size() == 1) {
            return 1;
        }

        int fleets = cars.size();

        System.out.println(cars);

        int cp = cars.size() - 1;
        for (int i = cars.size() - 2; i >= 0; i--) {
            if (cars.get(cp).timeToReach >= cars.get(i).timeToReach) {
                // makes a fleet
                fleets--;
            } else {
                cp = i;
            }
        }

        return fleets;

    }
}

class Car implements Comparable<Car> {
    int position;
    int speed;
    double timeToReach;

    public Car(int pos, int speed) {
        this.position = pos;
        this.speed = speed;
    }

    public int compareTo(Car c) {
        if (this.position < c.position) {
            return -1;
        } else if (this.position > c.position) {
            return 1;
        }
        return 0;
    }

    @Override
    public String toString() {
        return this.position + ": " + this.speed + "-" + this.timeToReach;
    }
}
